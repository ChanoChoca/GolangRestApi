package routes

import (
	"flashpage/app/controllers"
	"flashpage/app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUsuarioRoutes(m *mux.Router) {
	m.Handle("/usuarios", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateUsuario))).Methods("POST")
	m.Handle("/usuarios", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ListUsuarios))).Methods("GET")
	m.Handle("/usuarios/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetUsuario))).Methods("GET")
	m.Handle("/usuarios/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateUsuario))).Methods("PATCH")
	m.Handle("/usuarios/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeleteUsuario))).Methods("DELETE")
}