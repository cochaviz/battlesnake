package main

import (
	"battlesnake/pkg/handler"
	"log"
	"net/http"
	"os"
)

const ServerID = "BattlesnakeOfficial/starter-snake-go"

// Middleware

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

// Main Entrypoint

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", withServerID(handler.HandleIndex))
	http.HandleFunc("/start", withServerID(handler.HandleStart))
	http.HandleFunc("/move", withServerID(handler.HandleMove))
	http.HandleFunc("/end", withServerID(handler.HandleEnd))

	log.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
