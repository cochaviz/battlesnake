package internal

import (
	"battlesnake/internal/data"
	"battlesnake/internal/solver"
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"errors"
	"log"
	"os"
)

func Info() api.BattlesnakeInfoResponse {
	log.Println("INFO")

	return api.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "cochaviz",
		Color:      "#488A16",
		Head:       "beluga",
		Tail:       "curled",
	}
}

// This function is called everytime your Battlesnake is entered into a game.
func Start(state api.GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
func End(state api.GameState) {
	log.Printf("%s END\n\n", state.Game.ID)

	if os.Getenv("ANALYSIS") != "" {
		log.Println("Writing anaylsis...")
		utils.PlotMeasurements("test-"+state.Game.ID+".csv", "test-"+state.Game.ID+".png")
	}
}

// This function is called on every turn of a game. Use the provided GameState to decide
func Move(state api.GameState) api.BattlesnakeMoveResponse {
	nextMove, err := think(*data.ConvertFrom(state))

	if err != nil {
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)

		if os.Getenv("ANALYSIS") != "" {
			log.Println("Writing anaylsis...")
			utils.PlotMeasurements("test-"+state.Game.ID+".csv", "test-"+state.Game.ID+".png")
		}
	} else {
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)

		if os.Getenv("ANALYSIS") != "" {
			measurement := utils.Measurement{Turn: state.Turn, Length: int(state.You.Length)}
			measurement.AppendToFile("test-" + state.Game.ID + ".csv")
		}
	}
	return api.BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func think(state data.GameState) (string, error) {
	solution, _ := solver.Dfs(state, 11)

	if len(solution) == 0 {
		return "down", errors.New("Could not find a legal move")
	}
	return solution[0], nil
}
