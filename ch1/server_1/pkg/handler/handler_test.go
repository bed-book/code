package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../handler"
)

func TestGreet(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	handler.Greet(res, req)

	have := res.Body.String()
	want := "Hello, gophers!"

	if have != want {
		t.Fatalf(`incorrect response: have "%s", want "%s"`, have, want)
	}
}
