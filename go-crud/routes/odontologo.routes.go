package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhongonar/go-crud/db"
	"github.com/jhongonar/go-crud/models"
)

func GetOdontologosHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var odonto models.Odontologo
	db.DB.First(&odonto, params["id"])
	if odonto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Odontologo no encontrado"))
		return
	}
	json.NewEncoder(w).Encode(&odonto)
}

func PostOdontologosHandler(w http.ResponseWriter, r *http.Request) {
	var odonto models.Odontologo
	json.NewDecoder(r.Body).Decode(&odonto)
	createdOdontologo := db.DB.Create(&odonto)
	err := createdOdontologo.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&odonto)

}
func DeleteOdontologosHandler(w http.ResponseWriter, r *http.Request) {
	var odonto models.Odontologo
	params := mux.Vars(r)
	db.DB.First(&odonto, params["id"])
	if odonto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Odontologo no encontrado"))
		return
	}
	db.DB.Unscoped().Delete(&odonto)
	w.WriteHeader(http.StatusOK)

}
func PutOdontologosHandler(w http.ResponseWriter, r *http.Request) {
	var odonto models.Odontologo
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&odonto)
	updateOdontolog := db.DB.Where("id = ?", params["id"]).Updates(&odonto)
	err := updateOdontolog.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&odonto)
}
