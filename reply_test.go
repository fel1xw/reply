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

func TestCustomReplier(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	reply.Custom(w, http.StatusInternalServerError, nil)
	is.Equal(w.Code, http.StatusInternalServerError)
}

func TestDefaultMethods(t *testing.T) {
	is := is.New(t)
	replier := reply.NewReplier(reply.JSONMode)

	for _, c := range []struct {
		fn   func(w http.ResponseWriter, data interface{}) error
		code int
		data interface{}
	}{
		{
			fn:   replier.Ok,
			code: 200,
		},
		{
			fn:   replier.Success,
			code: 200,
		},
		{
			fn:   replier.Created,
			code: 201,
		},
		{
			fn:   replier.NotFound,
			code: 404,
		},
	} {
		w := httptest.NewRecorder()
		c.fn(w, c.data)
		resp := w.Result()
		is.Equal(resp.StatusCode, c.code)
		is.Equal(resp.Header.Get(reply.HeaderContentType), reply.MIMEApplicationJSON)
	}
}

func TestNewReplier(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	replier := reply.NewReplier(reply.XMLMode)
	replier.Ok(w, nil)
	resp := w.Result()
	is.Equal(resp.Header.Get(reply.HeaderContentType), reply.MIMEApplicationXML)
}

func TestSetHeader(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	replier := reply.NewReplier(reply.SetHeader(reply.HeaderLocation, "/test"))
	replier.Ok(w, nil)
	resp := w.Result()
	is.Equal(resp.Header.Get(reply.HeaderLocation), "/test")
}
