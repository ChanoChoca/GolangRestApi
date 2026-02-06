package routes

import (
	"flashpage/app/controllers"
	"flashpage/app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")

	r.Handle("/me",
		middlewares.AuthMiddleware(
			http.HandlerFunc(controllers.Me),
		),
	).Methods("GET")
}