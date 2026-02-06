package controllers

import (
	"encoding/json"
	"flashpage/app/config"
	"flashpage/app/middlewares"
	"flashpage/app/models/dto"
	"flashpage/app/services"
	"flashpage/app/utils"
	"fmt"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, ok := services.ValidateUser(req.DNI, req.Password)
    if !ok {
    	http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
    	return
    }

	token, _ := utils.GenerateJWT(int(user.ID))

	http.SetCookie(w, &http.Cookie{
		Name: "auth_token",
		Value: token,
		HttpOnly: true,
		Secure: config.CookieSecure,
		Path: "/",
		Domain: config.CookieDomain,
		Expires: time.Now().Add(time.Duration(config.JwtExpirationHours) * time.Hour),
		SameSite: config.CookieSameSite,
	})

	w.Write([]byte("login ok"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "auth_token",
		Value: "",
		HttpOnly: true,
		Secure: config.CookieSecure,
		Path: "/",
		Domain: config.CookieDomain,
		Expires: time.Unix(0, 0),
		SameSite: config.CookieSameSite,
	})

	w.Write([]byte("logout ok"))
}

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
	})
}

func User(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey)
	w.Write([]byte("usuario autenticado id=" + fmt.Sprint(userID)))
}