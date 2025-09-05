package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// Contoh login sederhana (tanpa DB, hanya hardcode)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Decode error:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Println("Login attempt:", req.Username, req.Password)

	if req.Username == "admin" && req.Password == "password123" {
		resp := LoginResponse{
			Message: "Login success",
			Token:   "mock-jwt-token-123", // nanti bisa ganti JWT beneran
		}
		json.NewEncoder(w).Encode(resp)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
