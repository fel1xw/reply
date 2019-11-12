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

// SetHeader -
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

// Ok -
func (r *Replier) Ok(w http.ResponseWriter, data interface{}) error {
	return r.Custom(w, http.StatusOK, data)
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

// Custom - Sets content-type to be application/json and encodes data + sets status code
func Custom(w http.ResponseWriter, statusCode int, data interface{}) error {
	return DefaultReplier.Custom(w, statusCode, data)
}

// Ok -
func Ok(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Custom(w, http.StatusOK, data)
}

// Success -
func Success(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Ok(w, data)
}

// Created -
func Created(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Custom(w, http.StatusCreated, data)
}

// NotFound -
func NotFound(w http.ResponseWriter, data interface{}) error {
	return DefaultReplier.Custom(w, http.StatusNotFound, data)
}
