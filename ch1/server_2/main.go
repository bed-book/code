package main

import (
	"log"
	"net/http"

	"github.com/techmexdev/back-end-dev-book/ch1/server_2/pkg/handler"
)

func main() {
	http.HandleFunc("/greet", handler.Greet)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
