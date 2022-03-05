package solver

import (
	"battlesnake/internal/data"
	"log"
	"math"
)

// Will try to find the longest path which is not terminal
func Dfs(state data.GameState, depth int) ([]string, int32) {
	if depth == 0 {
		cost := 0 * state.You.Length

		if state.IsTerminal() {
			cost = math.MaxInt32 / 2
		}
		return []string{}, cost
	}
	bestPath := []string{}
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
