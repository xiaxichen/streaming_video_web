package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"streaming_video_web/api/handler"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{
		r: r,
	}
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUser)
	router.POST("/user/:user_name", handler.Login)
	return router
}

func main() {
	router := RegisterHandler()
	wareHandler := NewMiddleWareHandler(router)
	http.ListenAndServe(":8000", wareHandler)
}
