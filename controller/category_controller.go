package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, response *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, response *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, response *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, response *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, response *http.Request, params httprouter.Params)
}