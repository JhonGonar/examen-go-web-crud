package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhongonar/go-crud/db"
	"github.com/jhongonar/go-crud/models"
)

func GetTurnoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var turno models.Turno
	db.DB.First(&turno, params["id"])
	if turno.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Turno no encontrado"))
		return
	}
	json.NewEncoder(w).Encode(&turno)
}

func PostTurnoHandler(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno
	json.NewDecoder(r.Body).Decode(&turno)
	createdOdontologo := db.DB.Create(&turno)
	err := createdOdontologo.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&turno)

}
func DeleteTurnoHandler(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno
	params := mux.Vars(r)
	db.DB.First(&turno, params["id"])
	if turno.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Turno no encontrado"))
		return
	}
	db.DB.Unscoped().Delete(&turno)
	w.WriteHeader(http.StatusOK)

}
func GetPacienteTurnoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var turno models.Turno
	var paciente models.Paciente
	db.DB.First(&turno, "paciente_id = ?", params["id"])
	if turno.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Turno no encontrado"))
		return
	}
	db.DB.Model(&turno).Association("PacienteID").Find(&paciente)
	//db.DB.Model(&turno).Association("PacienteId").Find(&paciente.ID)
	json.NewEncoder(w).Encode(&turno)
}
func PutTurnoHandler(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&turno)
	updateTurno := db.DB.Where("id = ?", params["id"]).Updates(&turno)
	err := updateTurno.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&turno)
}
