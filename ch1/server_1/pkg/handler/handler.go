package handler

import (
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, gophers!"))
}
