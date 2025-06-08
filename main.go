package main

import (
	"bankserver/auth"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.POST("/hello/:name", auth.CreateUser)
	log.Fatal(http.ListenAndServe(":8080", router))
}
