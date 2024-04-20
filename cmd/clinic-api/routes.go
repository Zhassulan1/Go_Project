package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// CLinic Singleton
	// Create a new appointment
	v1.HandleFunc("/appointments", app.createAppointmentHandler).Methods("POST")
	// Get a specific appointment
	v1.HandleFunc("/appointments/{appointmentId:[0-9]+}", app.getAppointmentHandler).Methods("GET")
	// Update a specific appointment
	v1.HandleFunc("/appointments/{appointmentId:[0-9]+}", app.updateAppointmentHandler).Methods("PUT")
	// Delete a specific appointment
	v1.HandleFunc("/appointments/{appointmentId:[0-9]+}", app.deleteAppointmentHandler).Methods("DELETE")

	// Create a new doctor
	v1.HandleFunc("/doctors", app.createDoctorHandler).Methods("POST")
	// Get a doctors list by pagination and filters
	v1.HandleFunc("/doctors", app.SearchDoctorHandler).Methods("GET")
	// Get a specific doctor
	v1.HandleFunc("/doctors/{doctorId:[0-9]+}", app.getDoctorHandler).Methods("GET")
	// Update a specific doctor
	v1.HandleFunc("/doctors/{doctorId:[0-9]+}", app.updateDoctorHandler).Methods("PUT")
	// Delete a specific doctor
	v1.HandleFunc("/doctors/{doctorId:[0-9]+}", app.deleteDoctorHandler).Methods("DELETE")

	// Create a new patient
	v1.HandleFunc("/patients", app.createPatientHandler).Methods("POST")
	// Get a specific patient
	v1.HandleFunc("/patients/{patientId:[0-9]+}", app.getPatientHandler).Methods("GET")
	// Update a specific patient
	v1.HandleFunc("/patients/{patientId:[0-9]+}", app.updatePatientHandler).Methods("PUT")
	// Delete a specific patient
	v1.HandleFunc("/patients/{patientId:[0-9]+}", app.deletePatientHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)

	return app.authenticate(r)
}
