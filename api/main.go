package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"streaming_video_web/api/handler"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUser)
	router.POST("/user/:user_name", handler.Login)
	return router
}

func main() {
	router := RegisterHandler()
	http.ListenAndServe(":8000", router)
}
