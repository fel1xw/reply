package reply_test

import (
	"encoding/xml"
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
		is.Equal(resp.Header.Get(reply.HeaderContentType), reply.MIMEApplicationJSON)
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
	replier := reply.NewReplier()

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

func TestCreatedWithLocation(t *testing.T) {
	is := is.New(t)
	userResourceURL := "/user/1"
	w := httptest.NewRecorder()
	replier := reply.NewReplier(reply.XMLMode)
	replier.CreatedWithLocation(w, userResourceURL, nil)
	resp := w.Result()
	is.Equal(resp.Header.Get(reply.HeaderLocation), userResourceURL)

	w = httptest.NewRecorder()
	reply.CreatedWithLocation(w, userResourceURL, nil)
	resp = w.Result()
	is.Equal(resp.Header.Get(reply.HeaderLocation), userResourceURL)
}

func TestNewReplier(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	replier := reply.NewReplier(reply.XMLMode)
	replier.Ok(w, nil)
	resp := w.Result()
	is.Equal(resp.Header.Get(reply.HeaderContentType), reply.MIMEApplicationXML)
}

func TestXMLModeError(t *testing.T) {
	type example struct {
		A int
		B int
		C func() int
	}
	is := is.New(t)
	w := httptest.NewRecorder()
	replier := reply.NewReplier(reply.XMLMode)
	err := replier.Ok(w, &example{})
	_, ok := err.(*xml.UnsupportedTypeError)

	is.True(ok)
}

func TestSetHeader(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	replier := reply.NewReplier(reply.SetHeader(reply.HeaderLocation, "/test"))
	replier.Ok(w, nil)
	resp := w.Result()
	is.Equal(resp.Header.Get(reply.HeaderLocation), "/test")
}

func TestSetHeaderFunc(t *testing.T) {
	is := is.New(t)
	w := httptest.NewRecorder()
	n := "5"
	r := reply.NewReplier(reply.SetHeaderFunc(func(w http.ResponseWriter) {
		w.Header().Set("test", n)
		w.Header().Set(reply.HeaderLocation, "/here")
	}))
	r.Ok(w, nil)
	resp := w.Result()
	is.Equal(resp.Header.Get("test"), n)
	is.Equal(resp.Header.Get(reply.HeaderLocation), "/here")
}
