package api

import (
	"net/http"
	"tugas-deployment/controller"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	SetupRoutes(router)
	router.ServeHTTP(w, r)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/belanja", controller.GetDaftarBelanja).Methods("GET")
	r.HandleFunc("/api/belanja", controller.CreateBelanja).Methods("POST")
	r.HandleFunc("/api/belanja/{id}", controller.GetBelanjaByID).Methods("GET")
}
