package handler

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := io.WriteString(writer, "Create User Handler")
	if err != nil {
		return
	}
}

func Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	user_name := params.ByName("user_name")
	_, err := io.WriteString(writer, user_name)
	if err != nil {
		return
	}
}
