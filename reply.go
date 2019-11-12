package reply

import (
	"encoding/json"
	"net/http"
)

// Library
const (
	Version = "0.0.1"
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
	Header http.Header
}

// Replier -
type Replier struct {
	config *ReplierConfig
}

var (
	// DefaultReplierConfig is the default configuration (json).
	DefaultReplierConfig = &ReplierConfig{
		Header: http.Header{
			HeaderContentType: []string{MIMEApplicationJSON},
		},
	}

	// DefaultReplier -
	DefaultReplier = NewReplier(JSONMode)
)

// Configure -
type Configure func(config *ReplierConfig)

// NewReplier creates a new replier with custom configuration
func NewReplier(config ...Configure) *Replier {
	cfg := &ReplierConfig{
		Header: http.Header{},
	}

	for _, configure := range config {
		configure(cfg)
	}

	return &Replier{
		config: cfg,
	}
}

// SetHeader sets a custom header to the Replier configuration
func SetHeader(key, value string) func(config *ReplierConfig) {
	return func(config *ReplierConfig) {
		config.Header.Set(key, value)
	}
}

// JSONMode - Sets content-type for responses to be content-type application/json
func JSONMode(config *ReplierConfig) {
	config.Header.Set(HeaderContentType, MIMEApplicationJSON)
}

// XMLMode - Sets content-type for responses to be content-type application/xml
func XMLMode(config *ReplierConfig) {
	config.Header.Set(HeaderContentType, MIMEApplicationXML)
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
	w.Header().Set(HeaderContentType, r.config.Header.Get(HeaderContentType))
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
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
