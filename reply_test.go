package reply_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fel1xw/reply"
	"github.com/matryer/is"
)

type testWriter struct {
	http.ResponseWriter
	statusCode int
}

func (t *testWriter) WriteHeader(statusCode int) {
	t.statusCode = statusCode

	t.WriteHeader(statusCode)
}

func TestReply(t *testing.T) {
	is := is.New(t)
	for _, c := range []struct {
		fn   func(w http.ResponseWriter, data interface{}) error
		code int
		data interface{}
		body string
	}{
		{
			fn:   reply.Ok,
			code: 200,
		},
		{
			fn:   reply.Success,
			code: 200,
		},
		{
			fn:   reply.Created,
			code: 201,
		},
		{
			fn:   reply.Created,
			code: 201,
		},
		{
			fn:   reply.NotFound,
			code: 404,
		},
	} {
		w := httptest.NewRecorder()
		c.fn(w, c.data)
		resp := w.Result()
		is.Equal(resp.StatusCode, c.code)
		is.Equal(resp.Header.Get("Content-Type"), "application/json")
	}
}
