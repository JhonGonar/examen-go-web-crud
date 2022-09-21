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

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/odontologo/{id}", routes.GetOdontologosHandler).Methods("GET")
	r.HandleFunc("/odontologo", routes.PostOdontologosHandler).Methods("POST")
	r.HandleFunc("/odontologo/{id}", routes.DeleteOdontologosHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
