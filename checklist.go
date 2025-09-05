package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Checklist struct {
	ID       int    `json:"id"`
	Category string `json:"category"` // machine / quality / environment
	Task     string `json:"task"`
	Status   string `json:"status"` // pending / done
}

var checklists = []Checklist{}
var nextID = 1

func ChecklistHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := json.NewEncoder(w).Encode(checklists)
		if err != nil {
			return
		}

	case http.MethodPost:
		var c Checklist
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		c.ID = nextID
		nextID++
		checklists = append(checklists, c)
		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			return
		}

	case http.MethodPut:
		var c Checklist
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		for i := range checklists {
			if checklists[i].ID == c.ID {
				checklists[i] = c
				err := json.NewEncoder(w).Encode(c)
				if err != nil {
					return
				}
				return
			}
		}
		http.Error(w, "Checklist not found", http.StatusNotFound)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}
		id, _ := strconv.Atoi(idStr)

		for i := range checklists {
			if checklists[i].ID == id {
				checklists = append(checklists[:i], checklists[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Checklist not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
