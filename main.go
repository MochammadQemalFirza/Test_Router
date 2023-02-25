package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	server := http.Server{
		Handler: router,
		Addr:    "localhost:8082",
	}
	server.ListenAndServe()
}
