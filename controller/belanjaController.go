package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tugas-deployment/models"

	"github.com/gorilla/mux"
)

func GetDaftarBelanja(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DaftarBelanja)
}

func CreateBelanja(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var belanja models.Belanja
	_ = json.NewDecoder(r.Body).Decode(&belanja)
	belanja.ID = len(models.DaftarBelanja) + 1
	models.DaftarBelanja = append(models.DaftarBelanja, belanja)
	json.NewEncoder(w).Encode(belanja)
}

func GetBelanjaByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, belanja := range models.DaftarBelanja {
		if belanja.ID == id {
			json.NewEncoder(w).Encode(belanja)
			return
		}
	}
	http.Error(w, "Belanja not found", http.StatusNotFound)
}
