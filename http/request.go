package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/priestess-dev/infra/v1/utils/misc"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
)

type ContentTypeEnum string

const (
	ContentTypeJson          ContentTypeEnum = "application/json"
	ContentTypeForm          ContentTypeEnum = "application/x-www-form-urlencoded"
	ContentTypeMultipartForm ContentTypeEnum = "multipart/form-data"
)

func (c ContentTypeEnum) String() string {
	return string(c)
}

// EmptyRequest is a empty request
type EmptyRequest struct{}

// RequestProcessor is a function that process request
type RequestProcessor func(r *http.Request, raw interface{}) error

func URLQueryProcessor[ReqT interface{}](r *http.Request, raw ReqT) {
	query := r.URL.Query()
	reqVal := reflect.ValueOf(raw)
	// panic if raw is not struct
	if reqVal.Kind() != reflect.Struct {
		panic(fmt.Errorf("raw must be struct"))
	}
	for i := 0; i < reqVal.NumField(); i++ {
		field := reqVal.Field(i)
		if field.IsValid() {
			// todo: unsafe if field is nested struct
			fn := misc.GetJsonName(reqVal.Type().Field(i).Tag.Get("json"), reqVal.Type().Field(i).Name)
			if fn != "" {
				query.Add(fn, fmt.Sprintf("%v", field.Interface()))
			}
		}
	}
	r.URL.RawQuery = query.Encode()
}

func JSONProcessor[ReqT interface{}](r *http.Request, raw ReqT) error {
	body, err := json.Marshal(raw)
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(bytes.NewReader(body))
	return nil
}

func MultipartFormRequestProcessor[ReqT interface{}](r *http.Request, raw ReqT, fileName string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	reqVal := reflect.ValueOf(raw)
	for i := 0; i < reqVal.NumField(); i++ {
		field := reqVal.Field(i)
		if field.IsValid() {
			if field.Type().Implements(reflect.TypeOf((*io.Reader)(nil)).Elem()) {
				// if the field is io.Reader, then add to multipart
				part, err := writer.CreateFormFile(misc.GetJsonName(reqVal.Type().Field(i).Tag.Get("json"), reqVal.Type().Field(i).Name), fileName)
				if err != nil {
					return err
				}
				//switch t := field.Interface().(type) {
				//case *os.File:
				//	_, err := io.Copy(part, t)
				//	if err != nil {
				//		return err
				//	}
				//case *bytes.Buffer:
				//	_, err := io.Copy(part, t)
				//	if err != nil {
				//		return err
				//	}
				//case io.Reader:
				//	_, err := io.Copy(part, t)
				//	if err != nil {
				//		return err
				//	}
				//default:
				//	return fmt.Errorf("unsupported type: %s", field.Type().String())
				//}
				_, err = io.Copy(part, field.Interface().(io.Reader))
				if err != nil {
					return err
				}
			} else {
				// else add to multipart form as key - value
				err := writer.WriteField(reqVal.Type().Field(i).Tag.Get("json"), fmt.Sprintf("%v", field.Interface()))
				if err != nil {
					return err
				}
			}
		}
	}
	err := writer.Close()
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	return nil
}

func EndpointHandler[ReqT, RespT interface{}](client *http.Client, url string, method string, headers map[string]string, multipartName string) func(ReqT) (RespT, error) {
	return func(req ReqT) (RespT, error) {
		reqVal := reflect.ValueOf(req)
		var respObj RespT
		if reqVal.Kind() != reflect.Struct {
			panic(fmt.Errorf("raw must be struct"))
		}
		httpReq, err := http.NewRequest(method, url, nil)
		if multipartName != "" {
			err = MultipartFormRequestProcessor(httpReq, req, multipartName)
			if err != nil {
				// build empty RespT
				return respObj, err
				//return , err
			}
		} else {
			switch method {
			case http.MethodGet:
				URLQueryProcessor(httpReq, req)
			case http.MethodPost:
				fallthrough
			case http.MethodPatch:
				err = JSONProcessor(httpReq, req)
				if err != nil {
					return respObj, err
				}
			default:
				return respObj, fmt.Errorf("unsupported method: %s", method)
			}
		}
		for k, v := range headers {
			httpReq.Header.Add(k, v)
		}
		fmt.Printf("[%s] to %s\n", httpReq.Method, httpReq.URL.String())
		resp, err := client.Do(httpReq)
		if err != nil {
			return respObj, err
		} else if resp.StatusCode != http.StatusOK {
			// todo: log
			return respObj, fmt.Errorf("failed to do request, status code: %d", resp.StatusCode)
		} else {
			err := json.NewDecoder(resp.Body).Decode(&respObj)
			if err != nil {
				return respObj, err
			}
			return respObj, nil
		}
	}
}
