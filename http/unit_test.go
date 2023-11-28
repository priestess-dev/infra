package http

import (
	"net/http"
	"testing"
)

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
