package main

import (
	"bankserver/auth"
	"bankserver/db"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	err = db.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
	}

	router := httprouter.New()
	router.POST("/create-user/:name", auth.CreateUser)
	log.Fatal(http.ListenAndServe(":8005", router))
}
