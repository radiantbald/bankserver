package main

import (
	"bankserver/auth"
	"bankserver/db"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	err := db.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
	}

	router := httprouter.New()
	router.POST("/create-user/:name", auth.CreateUser)
	log.Fatal(http.ListenAndServe(":8080", router))
}
