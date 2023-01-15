package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosGenuino/api-rest-go/database"
	"github.com/CarlosGenuino/api-rest-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page Formatado")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(&personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(&personalidade)
}