package controllers

import (
	"encoding/json"
	"flashpage/app/models"
	"flashpage/app/models/dto"
	"flashpage/app/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUsuario(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateUsuarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json inválido", 400)
		return
	}

	u := models.Usuario{
		Nombre: req.Nombre,
		Apellido: req.Apellido,
		DNI: req.DNI,
		Telefono: req.Telefono,
		Rol: models.Rol(req.Rol),
		SupervisorID: req.SupervisorID,
		Provincia: req.Provincia,
		Localidad: req.Localidad,
		Direccion: req.Direccion,
		CodigoPostal: req.CodigoPostal,
		Password: req.Password,
	}

	if err := services.CreateUsuario(&u); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "id inválido", 400)
		return
	}

	u, err := services.GetUsuario(uint(id))
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func ListUsuarios(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("pageSize"))

	result, err := services.ListUsuarios(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {

	var req dto.UpdateUsuarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json inválido", 400)
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "id inválido", 400)
		return
	}

	u := models.Usuario{
		ID: uint(id),
		Nombre: req.Nombre,
		Apellido: req.Apellido,
		Telefono: req.Telefono,
		SupervisorID: req.SupervisorID,
		Provincia: req.Provincia,
		Localidad: req.Localidad,
		Direccion: req.Direccion,
		CodigoPostal: req.CodigoPostal,
		Activo: req.Activo,
	}

	if err := services.UpdateUsuario(&u); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("updated"))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", 400)
		return
	}

	if err := services.DisableUsuario(uint(id)); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("disabled"))
}