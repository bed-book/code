package main

import (
	"log"
	"net/http"

	"github.com/techmexdev/back-end-dev-book/ch1/server_1/pkg/handler"
)

func main() {
	http.HandleFunc("/", handler.Greet)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
