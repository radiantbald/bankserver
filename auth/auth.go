package auth

import (
	"bankserver/db"
	"bankserver/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type registerRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req registerRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Surname == "" || req.PhoneNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:        req.Name,
		Surname:     req.Surname,
		PhoneNumber: req.PhoneNumber,
		Verified:    false,
	}
	
	db.CreateUser(&user)

	fmt.Fprintf(w, "Пользователь создан, %s!\n", req.Name)
}
