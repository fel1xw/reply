![Logo](https://i.ibb.co/k3sjyR3/reply-fel1xw.png)

# Reply [![GoDoc](https://godoc.org/github.com/fel1xw/reply?status.png)](http://godoc.org/github.com/fel1xw/reply) [![Build Status](https://github.com/fel1xw/reply/workflows/Go/badge.svg)](https://github.com/fel1xw/reply/workflows/Go/badge.svg) [![Go Version](https://img.shields.io/github/go-mod/go-version/fel1xw/reply)](https://img.shields.io/github/go-mod/go-version/fel1xw/reply) [![Go Report Card](https://goreportcard.com/badge/github.com/fel1xw/reply)](https://goreportcard.com/report/github.com/fel1xw/reply)

* Easy to write and read
* JSON mode enabled by default
  * Sets `Content-Type` to be `application/json`
  * uses `json.NewDecoder(w).Encode(data)`
* Simplifies status code responses

### Usage

```go
type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
  u := &User{ID: "123", Name: "Felix"}
	reply.Created(w, u)
}
```
