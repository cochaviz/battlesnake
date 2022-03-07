package solver

import (
	"battlesnake/internal/data"
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"log"
	"math"
)

func maxDistanceFromOthers(state data.GameState) int32 {
	if state.IsTerminal() {
		return math.MaxInt32 - int32(state.Turn)
	}
	totalDistance := int32(0)

	for _, snake := range state.OtherSnakes {
		totalDistance += int32(utils.ManHattanDistance(state.You.Head, snake.Head))
	}
	return totalDistance - int32(state.Turn) + state.You.Health + state.You.Length
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
	return -state.You.Length - int32(state.Turn) + state.You.Health
}

// Will try to find the longest path which is not terminal
func Dfs(state data.GameState, depth int) ([]string, int32) {
	bestPath := []string{}

	if depth == 0 || state.IsTerminal() {
		return bestPath, maxDistanceFromOthers(state)
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
