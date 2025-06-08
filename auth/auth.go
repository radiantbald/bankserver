package auth

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Пользователь создан, %s!\n", ps.ByName("name"))
}
