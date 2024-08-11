package main

import (
	"log"
	"net/http"

	"github.com/marciomarinho/school-management-api-go/config"
	"github.com/marciomarinho/school-management-api-go/handlers"
	"github.com/marciomarinho/school-management-api-go/service"

	"github.com/gorilla/mux"
)

func main() {
	db := config.SetupDynamoDB()

	studentService := service.NewStudentService(db)
	studentHandler := &handlers.StudentHandler{Service: studentService}

	r := mux.NewRouter()
	r.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
