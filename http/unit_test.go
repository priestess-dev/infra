package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/priestess-dev/infra/utils/random"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestGetJsonTag(t *testing.T) {
	type Test struct {
		Test string `json:"test1,omitempty"`
	}
	tt := Test{
		Test: "test",
	}
	ttVal := reflect.ValueOf(tt)
	for i := 0; i < ttVal.NumField(); i++ {
		field := ttVal.Type().Field(i)
		t.Logf("field: %s, get tag: %s, ", field.Name, field.Tag.Get("json"))
		if tag, ok := field.Tag.Lookup("json"); ok {
			t.Logf("lookup tag: %s\n", tag)
		}
	}
}

func TestURLQueryProcessor(t *testing.T) {
	t.Run("Non Empty Req", func(t *testing.T) {
		type ReqType struct {
			Test string `json:"test"`
		}
		r, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api", nil)
		URLQueryProcessor[ReqType](r, ReqType{
			Test: "test",
		})
		t.Log(r.URL)
	})
	t.Run("Empty Req", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api", nil)
		URLQueryProcessor[EmptyRequest](r, EmptyRequest{})
		t.Log(r.URL)
	})
}

func TestJSONProcessor(t *testing.T) {
	t.Run("Non Empty Req", func(t *testing.T) {
		type ReqType struct {
			Test string `json:"test"`
		}
		r, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api", nil)
		err := JSONProcessor[ReqType](r, ReqType{
			Test: "test",
		})
		if err != nil {
			t.Fatal(err)
		}
		s := make([]byte, 1024)
		l, _ := r.Body.Read(s)
		t.Logf("Content: %s", string(s[:l]))
	})
	t.Run("Empty Req", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api", nil)
		err := JSONProcessor[EmptyRequest](r, EmptyRequest{})
		if err != nil {
			t.Fatal(err)
		}
		s := make([]byte, 1024)
		l, _ := r.Body.Read(s)
		t.Logf("Content: %s", string(s[:l]))
	})
}

func TestEndpointHandler(t *testing.T) {
	var FILE_STRING = "content of some file"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// start server
	go func() {
		type R struct {
			Test string `json:"test"`
		}
		server := NewServer("localhost", 8080)
		server.AddRoutes(EndpointConfig{
			Path:   "/api",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"test": random.RandString(10),
				})
			},
		}, EndpointConfig{
			Path:   "/api",
			Method: http.MethodPost,
			Handler: func(c *gin.Context) {
				// get params
				var req R
				err := c.ShouldBindJSON(&req)
				if err != nil {
					t.Error(err)
					return
				}
				for i, p := range c.Params {
					t.Logf("[SERVER] Param %d: %s - %s\n", i, p.Key, p.Value)
				}
				c.JSON(http.StatusOK, gin.H{
					"test": random.RandString(10),
				})
			},
		}, EndpointConfig{
			Path:   "/file",
			Method: http.MethodPost,
			Handler: func(c *gin.Context) {
				file, err := c.FormFile("test")
				t.Logf("[SERVER] File: %s, %d\n", file.Filename, file.Size)
				if err != nil {
					t.Error(err)
					return
				}
				// read file content
				f, err := file.Open()
				if err != nil {
					t.Error(err)
					return
				}
				defer f.Close()
				// read file content
				buf := make([]byte, file.Size)
				_, err = f.Read(buf)
				if err != nil {
					t.Error(err)
					return
				}
				f.Read(buf)
				t.Logf("%s\n\n", string(buf))
				if string(buf) != FILE_STRING {
					t.Errorf("file content is not equal")
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"test": random.RandString(10),
				})
			},
		})
		err := server.Start()
		if err != nil {
			return
		}
	}()
	select {
	case <-ctx.Done():
		t.Log("server might started")
	}
	t.Run("GET", func(t *testing.T) {
		type ReqType struct {
			Test string `json:"test"`
		}
		type RespType struct {
			Test string `json:"test"`
		}
		handler := EndpointHandler[ReqType, RespType](http.DefaultClient, "http://localhost:8080/api", http.MethodGet, nil, "")
		resp, err := handler(ReqType{
			Test: "test",
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("[CLIENT]: %s\n", resp.Test)
	})
	t.Run("POST", func(t *testing.T) {
		type ReqType struct {
			Test string `json:"test"`
		}
		type RespType struct {
			Test string `json:"test"`
		}
		handler := EndpointHandler[ReqType, RespType](http.DefaultClient, "http://localhost:8080/api", http.MethodPost, nil, "")
		resp, err := handler(ReqType{
			Test: "test",
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("[CLIENT]: %s\n", resp.Test)
	})
	t.Run("POST with multipart", func(t *testing.T) {
		type ReqType struct {
			Test io.Reader `json:"test"`
		}
		type RespType struct {
			Test string `json:"test"`
		}
		handler := EndpointHandler[ReqType, RespType](http.DefaultClient, "http://localhost:8080/file", http.MethodPost, nil, "test_ng.txt")
		//f, err := os.OpenFile("tmp/test.txt", os.O_RDONLY, 0644)
		f := strings.NewReader(FILE_STRING)
		//fstat, _ := f.Stat()
		buf := make([]byte, f.Size())
		f.Read(buf)
		t.Logf("[CLIENT] File: %s\n", string(buf))

		//defer f.Close()
		_, err := f.Seek(0, io.SeekStart)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := handler(ReqType{
			Test: f,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("[CLIENT]: %s\n", resp.Test)
	})
}
