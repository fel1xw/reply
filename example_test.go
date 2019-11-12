package reply_test

import (
	"net/http/httptest"

	"github.com/fel1xw/reply"
)

func ExampleOk() {
	w := httptest.NewRecorder()
	reply.Ok(w, nil)

	// Output: statusCode set to 200
}
