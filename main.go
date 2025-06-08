package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type requestData struct {
	Name string `json:"name"`
}

type responseData struct {
	Name string `json:"greeting"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Worlddddd")
}

func jsonHelloWorld(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var data requestData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return
		}
		var response = responseData{
			Name: "Hello " + data.Name,
		}
		json.NewEncoder(w).Encode(response)
	case http.MethodGet:
		var name = r.URL.Query().Get("name")
		var response = responseData{
			Name: "GET " + name,
		}
		json.NewEncoder(w).Encode(response)
	}

}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/json", jsonHelloWorld)

	http.ListenAndServe(":8080", nil)
}
