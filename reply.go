package reply

import (
	"encoding/json"
	"net/http"
)

// Custom - Sets content-type to be application/json and encodes data + sets status code
func Custom(w http.ResponseWriter, statusCode int, data ...interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

// Ok -
func Ok(w http.ResponseWriter, data ...interface{}) error {
	return Custom(w, http.StatusOK, data)
}

// Success -
func Success(w http.ResponseWriter, data ...interface{}) error {
	return Ok(w, data)
}

// Created -
func Created(w http.ResponseWriter, data ...interface{}) error {
	return Custom(w, http.StatusCreated, data)
}

// NotFound -
func NotFound(w http.ResponseWriter, data ...interface{}) error {
	return Custom(w, http.StatusNotFound, data)
}
