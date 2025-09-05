package main

import (
	"net/http"

	"github.com/Dimas263/backend-checklist-system-manufacture/api"
	"github.com/vercel/go-bridge/go/bridge"
)

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*w.Header().Set("Access-Control-Allow-Origin", "*")*/
		w.Header().Set("Access-Control-Allow-Origin", "https://checklist-system-manufacture.vercel.app")
		/*w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")*/
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// kalau OPTIONS (preflight), balas kosong
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

	bridge.Start(mux) // <-- masukkan mux sebagai http.Handler
}
