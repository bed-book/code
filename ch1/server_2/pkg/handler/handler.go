package handler

import (
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Write([]byte("Hello, " + name + "!"))
}
