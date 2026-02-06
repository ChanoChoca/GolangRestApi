package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	RegisterAuthRoutes(r)

	RegisterUsuarioRoutes(r)

	return r
}