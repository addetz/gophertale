package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/addetz/gophertale/tdd/data"
)

type Handler struct {
	es *data.EmployeeService
}

func NewHandler(es *data.EmployeeService) *Handler {
	return &Handler{
		es: es,
	}
}

func (h *Handler) UpsertEmployee(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := readRequestBody(r)
	// Handle any errors & write an error HTTP status & response
	if err != nil {
		http.Error(w, "error reading body", http.StatusBadRequest)
		return
	}

	// Initialize a user to unmarshal request body into
	var employee data.Employee
	if err := json.Unmarshal(body, &employee); err != nil {
		http.Error(w, "invalid user body", http.StatusBadRequest)
		return
	}

	// Call the repository method corresponding to the operation
	_, err = h.es.Add(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send an HTTP success status & the return value from the repo
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func (h *Handler) ListEmployeeByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "no id provided", http.StatusBadRequest)
		return
	}
	employee, err := h.es.Get(id)
	if err != nil {
		http.Error(w, "invalid employee id", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

// readRequestBody is a helper method that
// allows to read a request body and return any errors.
func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return []byte{}, err
	}
	if err := r.Body.Close(); err != nil {
		return []byte{}, err
	}
	return body, err
}
