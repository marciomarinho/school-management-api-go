package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marciomarinho/school-management-api-go/service"
	"github.com/marciomarinho/school-management-api-go/types"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	Service service.StudentService
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student types.Student
	_ = json.NewDecoder(r.Body).Decode(&student)

	err := h.Service.CreateStudent(r.Context(), &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	student, err := h.Service.GetStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if student == nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var student types.Student
	_ = json.NewDecoder(r.Body).Decode(&student)

	err := h.Service.UpdateStudent(r.Context(), id, &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := h.Service.DeleteStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
