package main

import (
	"battlesnake/pkg/handler"
	"log"
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

const ServerID = "BattlesnakeOfficial/starter-snake-go"

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

func initRelicAgent(key string) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Battlesnake Performance"),
		newrelic.ConfigLicense(key),
	)
	if err != nil {
		log.Panic("Could not initialize New Relic agent, please check the config")
	} else {
		log.Print("Succesfully initialized New Relic agent")
	}
	return app
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	newrelicKey := os.Getenv("NEWRELICKEY")

	// if there is a relic key in current environment
	if newrelicKey != "" {
		// init relic agent
		app := initRelicAgent(newrelicKey)
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/", withServerID(handler.HandleIndex)))
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/start", withServerID(handler.HandleStart)))
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/move", withServerID(handler.HandleMoveRelic)))
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/end", withServerID(handler.HandleEnd)))
	} else {
		// otherwise, don't bother
		http.HandleFunc("/", withServerID(handler.HandleIndex))
		http.HandleFunc("/start", withServerID(handler.HandleStart))
		http.HandleFunc("/move", withServerID(handler.HandleMove))
		http.HandleFunc("/end", withServerID(handler.HandleEnd))
	}
	log.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
