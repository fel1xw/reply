### Install

```bash
go get github.com/fel1xw/reply
```

### Use

```go
func CreatedHandler(w http.ResponseWriter, r *http.Request) {
  type user struct {
    ID string `json:"id"`
    Name string `json:"name"`
  }
  u := &user{"1", "Felix"}

  reply.Created(w, u)
}
```
