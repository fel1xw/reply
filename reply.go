package reply

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Headers
const (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderLocation      = "Location"
)

// Mime types
const (
	MIMEApplicationJSON = "application/json"
	MIMEApplicationXML  = "application/xml"
)

// ReplierConfig -
type ReplierConfig struct {
	Header      http.Header
	transformFn []func(w http.ResponseWriter)
	encode      func(w http.ResponseWriter, data interface{}) error
}

// Replier -
type Replier struct {
	config *ReplierConfig
}

var (
	// DefaultReplier -
	DefaultReplier = NewReplier()
)

// Configure -
type Configure func(config *ReplierConfig)

// NewReplier creates a new replier with custom configuration
func NewReplier(config ...Configure) *Replier {
	cfg := &ReplierConfig{
		Header: http.Header{},
	}

	JSONMode(cfg)

	for _, configure := range config {
		configure(cfg)
	}

	return &Replier{
		config: cfg,
	}
}

// SetHeader sets a custom header to the Replier configuration
func SetHeader(key, value string) Configure {
	return func(config *ReplierConfig) {
		config.Header.Set(key, value)
	}
}

// SetHeaderFunc is executed on each response call
func SetHeaderFunc(fn func(w http.ResponseWriter)) Configure {
	return func(config *ReplierConfig) {
		config.transformFn = append(config.transformFn, fn)
	}
}

// JSONMode - Sets content-type for responses to be content-type application/json
func JSONMode(config *ReplierConfig) {
	config.Header.Set(HeaderContentType, MIMEApplicationJSON)
	config.encode = func(w http.ResponseWriter, data interface{}) error {
		return json.NewEncoder(w).Encode(data)
	}
}

// XMLMode - Sets content-type for responses to be content-type application/xml
func XMLMode(config *ReplierConfig) {
	config.Header.Set(HeaderContentType, MIMEApplicationXML)
	config.encode = func(w http.ResponseWriter, data interface{}) error {
		x, err := xml.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		_, err = w.Write(x)
		return err
	}
}

// Ok sets the statusCode of the response to be 200
func (r *Replier) Ok(w http.ResponseWriter, data interface{}) error {
	return r.Custom(w, http.StatusOK, data)
}

// Success sets the statusCode of the response to be 200
func (r *Replier) Success(w http.ResponseWriter, data interface{}) error {
	return r.Custom(w, http.StatusOK, data)
}

// NotFound sets the statusCode of the response to be 404
func (r *Replier) NotFound(w http.ResponseWriter, data interface{}) error {
	return r.Custom(w, http.StatusNotFound, data)
}

// Created sets the statusCode of the response to be 201
func (r *Replier) Created(w http.ResponseWriter, data interface{}) error {
	return r.Custom(w, http.StatusCreated, data)
}

// Custom -
func (r *Replier) Custom(w http.ResponseWriter, statusCode int, data interface{}) error {
	for key, value := range r.config.Header {
		w.Header().Set(key, value[0])
	}
	for _, fn := range r.config.transformFn {
		fn(w)
	}
	w.WriteHeader(statusCode)

	return r.config.encode(w, data)
}

// Custom sets the statusCode of the response to the passed one with the DefaultConfig
func Custom(w http.ResponseWriter, statusCode int, data interface{}) error {
	return DefaultReplier.Custom(w, statusCode, data)
}

// Ok sets the statusCode of the response to be 200 with the DefaultConfig
func Ok(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Custom(w, http.StatusOK, data)
}

// Success sets the statusCode of the response to be 200 with the DefaultConfig
func Success(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Ok(w, data)
}

// Created sets the statusCode of the response to be 201 with the DefaultConfig
func Created(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Custom(w, http.StatusCreated, data)
}

// NotFound sets the statusCode of the response to be 404 with the DefaultConfig
func NotFound(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Custom(w, http.StatusNotFound, data)
}
