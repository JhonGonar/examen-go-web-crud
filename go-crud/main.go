package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhongonar/go-crud/db"
	"github.com/jhongonar/go-crud/models"
	"github.com/jhongonar/go-crud/routes"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola API"))
}
func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Odontologo{})
	db.DB.AutoMigrate(models.Paciente{})
	db.DB.AutoMigrate(models.Turno{})

	r := mux.NewRouter()
	//Odontologo

	r.HandleFunc("/odontologo/{id}", routes.GetOdontologosHandler).Methods("GET")
	r.HandleFunc("/odontologo", routes.PostOdontologosHandler).Methods("POST")
	r.HandleFunc("/odontologo/{id}", routes.DeleteOdontologosHandler).Methods("DELETE")
	r.HandleFunc("/odontologo/{id}", routes.PutOdontologosHandler).Methods("PUT")

	//Paciente
	r.HandleFunc("/paciente/{id}", routes.GetPacienteHandler).Methods("GET")
	r.HandleFunc("/paciente", routes.PostPacienteHandler).Methods("POST")
	r.HandleFunc("/paciente/{id}", routes.DeletePacienteHandler).Methods("DELETE")
	r.HandleFunc("/paciente/{id}", routes.PutPacienteHandler).Methods("PUT")

	//Turnos
	r.HandleFunc("/turno/{id}", routes.GetTurnoHandler).Methods("GET")
	r.HandleFunc("/turno", routes.PostTurnoHandler).Methods("POST")
	r.HandleFunc("/turno/{id}", routes.DeleteTurnoHandler).Methods("DELETE")
	r.HandleFunc("/turno/paciente/{id}", routes.GetPacienteTurnoHandler).Methods("GET")
	r.HandleFunc("/turno/{id}", routes.PutTurnoHandler).Methods("PUT")

	http.ListenAndServe(":3000", r)
}
