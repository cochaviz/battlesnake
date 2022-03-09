package handler

import (
	"battlesnake/internal"
	"battlesnake/pkg/api"
	"encoding/json"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"net/http"
)

// HTTP Handlers

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := internal.Info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	state := api.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}

	internal.Start(state)

	// Nothing to respond with here
}

func updateRelic(state api.GameState, r *http.Request) {
	// Update relic
	tnx := newrelic.FromContext(r.Context())
	customAttributes, err := api.GetCustomAttributesFromGameState(state)

	if err != nil {
		log.Printf("ERROR: Failed to convert GameState to relic CustomAttribute, %s", err)
	}
	for key, value := range customAttributes {
		tnx.AddAttribute(key, value)
	}
}

func HandleMoveRelic(w http.ResponseWriter, r *http.Request) {

	state := api.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}
	updateRelic(state, r)

	response := internal.Move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	state := api.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := internal.Move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func HandleEnd(w http.ResponseWriter, r *http.Request) {
	state := api.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}

	internal.End(state)

	// Nothing to respond with here
}
