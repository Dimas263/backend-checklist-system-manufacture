package main

import (
	_ "encoding/json"
	"net/http"

	"github.com/Dimas263/backend-checklist-system-manufacture/api"
	"github.com/vercel/go-bridge/go/bridge"
)

// Middleware CORS
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Izinkan domain frontend
		w.Header().Set("Access-Control-Allow-Origin", "https://frontend-checklist-system-manufactu.vercel.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", withCORS(api.LoginHandler))
	mux.HandleFunc("/checklist", withCORS(api.ChecklistHandler))

	bridge.Start(mux)
}
