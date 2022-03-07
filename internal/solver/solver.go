package solver

import (
	"battlesnake/internal/data"
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"log"
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

func minLength(state data.GameState) int32 {
	if state.IsTerminal() {
		return math.MaxInt32 - int32(state.Turn)
	}
	return -int32(state.Turn) + state.You.Health + state.You.Length
}

func maxLength(state data.GameState) int32 {
	if state.IsTerminal() {
		return math.MaxInt32 - int32(state.Turn)
	}
	return -state.You.Length - int32(state.Turn)
}

// Will try to find the longest path which is not terminal
func Dfs(state data.GameState, depth int) ([]string, int32) {
	bestPath := []string{}

	if depth == 0 || state.IsTerminal() {
		return bestPath, maxLength(state)
	}
	var lowestCost int32
	lowestCost = math.MaxInt32

	for _, move := range state.SafeMoves() {
		nextState, err := state.Move(move)

		if err != nil {
			log.Print("Cannot move from current state")
			return []string{}, math.MaxInt32 / 2
		}
		path, cost := Dfs(*nextState, depth-1)

		if cost < lowestCost {
			bestPath = append([]string{move}, path...)
			lowestCost = cost
		}
	}
	return bestPath, lowestCost
}
