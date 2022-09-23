package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhongonar/go-crud/db"
	"github.com/jhongonar/go-crud/models"
)

func GetPacienteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var paciente models.Paciente
	db.DB.First(&paciente, params["id"])
	if paciente.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Paciente no encontrado"))
		return
	}
	json.NewEncoder(w).Encode(&paciente)
}

func PostPacienteHandler(w http.ResponseWriter, r *http.Request) {
	var paciente models.Paciente
	json.NewDecoder(r.Body).Decode(&paciente)
	createdOdontologo := db.DB.Create(&paciente)
	err := createdOdontologo.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&paciente)

}
func DeletePacienteHandler(w http.ResponseWriter, r *http.Request) {
	var paciente models.Paciente
	params := mux.Vars(r)
	db.DB.First(&paciente, params["id"])
	if paciente.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Paciente no encontrado"))
		return
	}
	db.DB.Unscoped().Delete(&paciente)
	w.WriteHeader(http.StatusOK)

}
func PutPacienteHandler(w http.ResponseWriter, r *http.Request) {
	var paciente models.Paciente
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&paciente)
	updatePaciente := db.DB.Where("id = ?", params["id"]).Updates(&paciente)
	err := updatePaciente.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&paciente)
}
