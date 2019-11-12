### Install

```bash
go get -u github.com/fel1xw/reply
```

### Use

```go
func CreatedHandler(w http.ResponseWriter, r *http.Request) {
  type User struct {
    ID string `json:"id"`
    Name string `json:"name"`
  }
  u := &User{"1", "Felix"}
	reply.Created(w, u)
}
```
