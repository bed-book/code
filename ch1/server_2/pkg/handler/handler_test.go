package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/techmexdev/back-end-dev-book/ch1/server_2/pkg/handler"
)

func TestGreet(t *testing.T) {
	tt := []struct{ name string }{
		{name: "John"},
		{name: "Paul"},
		{name: "George"},
		{name: "Ringo"},
	}

	for _, tc := range tt {
		t.Run("Greet "+tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/greet?name="+tc.name, nil)
			if err != nil {
				t.Fatal(err)
			}

			res := httptest.NewRecorder()

			handler.Greet(res, req)

			have := res.Body.String()
			want := "Hello, " + tc.name + "!"

			if have != want {
				t.Fatalf(`incorrect response: have "%s", want "%s"`, have, want)
			}
		})
	}
}
