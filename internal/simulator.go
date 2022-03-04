package internal

import (
	"battlesnake/pkg/api"
)

func Rotate_you(state api.GameState) {
	if len(state.Board.Snakes) > 0 {
		return
	}
	you := state.You
	state.You = state.Board.Snakes[0]

	for index := range state.Board.Snakes[0 : len(state.Board.Snakes)-1] {
		state.Board.Snakes[index] = state.Board.Snakes[index+1]
	}
	state.Board.Snakes[len(state.Board.Snakes)-1] = you
}

func Simulate_others(state api.GameState) []string {
	moves := []string{}

	for range state.Board.Snakes {
		Rotate_you(state)
		moves = append(moves, Think(state))
	}
	return moves
}
