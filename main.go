package main

import (
	"github.com/Dimas263/backend-checklist-system-manufacture/api"
	"github.com/vercel/go-bridge/go/bridge"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", api.LoginHandler)
	mux.HandleFunc("/checklist", api.ChecklistHandler)

	bridge.Start(mux) // <-- masukkan mux sebagai http.Handler
}
