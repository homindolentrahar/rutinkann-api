package controller

import "net/http"

type AuthController interface {
	SignIn(writer http.ResponseWriter, request *http.Request)
	Register(writer http.ResponseWriter, request *http.Request)
}
