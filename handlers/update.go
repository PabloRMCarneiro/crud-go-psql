package handlers

import (
	"crud/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"fmt"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error converting id: %d to int", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Println("Error decoding JSON", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	
	if err != nil {
		log.Printf("Error updating todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Updated %d rows", rows)
	}

	resp := map[string]any{
		"Error": false,
		"Message": fmt.Sprintf("Todo updated with id: %d", id),
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
