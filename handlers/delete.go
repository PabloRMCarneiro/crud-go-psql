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

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error converting id: %d to int", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	
	if err != nil {
		log.Printf("Error remove todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Removed %d rows", rows)
	}

	resp := map[string]any{
		"Error": false,
		"Message": fmt.Sprintf("Todo removed with id: %d", id),
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
