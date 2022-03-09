package internal

import (
	"battlesnake/internal/data"
	"battlesnake/internal/solver"
	"battlesnake/pkg/api"
	"errors"
	"log"
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
}

// This function is called on every turn of a game. Use the provided GameState to decide
func Move(state api.GameState) api.BattlesnakeMoveResponse {
	nextMove, err := think(*data.ConvertFrom(state))

	if err != nil {
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)

	} else {
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return api.BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func think(state data.GameState) (string, error) {
	solution, _ := solver.Dfs(state, 10)
	log.Print(state.You.Length)

	if len(solution) == 0 {
		return "down", errors.New("Could not find a legal move")
	}
	return solution[0], nil
}
