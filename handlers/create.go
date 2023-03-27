package handlers

import (
	"crud/models"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	_"github.com/golang/protobuf/ptypes/any"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	
	if err != nil {
		log.Println("Error decoding JSON", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Error inserting todo: %s", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("Todo inserted with id: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}