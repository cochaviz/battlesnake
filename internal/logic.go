package internal

import (
	"battlesnake/internal/data"
	"battlesnake/internal/solver"
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"errors"
	"math"
)

// This function removes moves that will get the snake further away from food if there are other options
func greedyMove(state data.GameState) []string {
	head := state.You.Head
	possibleMoves := state.PossibleMoves()

	minDistance := math.MaxInt
	minFood := api.Coord{}

	for _, food := range state.Board.Food {
		var newDistance = utils.ManHattanDistance(head, food)

		if newDistance < minDistance {
			minFood = food
			minDistance = newDistance
		}
	}

	if minDistance != math.MaxInt {
		if minFood.X >= head.X {
			if utils.TryRemoveMove(possibleMoves, "left") {
			}
		}
		if minFood.X <= head.X {
			if utils.TryRemoveMove(possibleMoves, "right") {
			}
		}
		if minFood.Y >= head.Y {
			if utils.TryRemoveMove(possibleMoves, "down") {
			}
		}
		if minFood.Y <= head.Y {
			if utils.TryRemoveMove(possibleMoves, "up") {
			}
		}
	}
	return utils.SafeMoves(possibleMoves)
}

func Think(state data.GameState) (string, error) {
	solution, _ := solver.Dfs(state, 12)

	if len(solution) == 0 {
		return "down", errors.New("Could not find a legal move")
	}
	return solution[0], nil
}
