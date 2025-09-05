package api

import (
	"encoding/json"
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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

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

// Handler Fungsi yang diexport ke Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	LoginHandler(w, r)
}
