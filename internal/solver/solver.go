package solver

import (
	"battlesnake/internal/data"
	"log"
)

// Will try to find the longest path which is not terminal
func Dfs(state data.GameState, depth int) ([]string, bool) {
	if depth == 0 {
		return []string{}, false
	}
	longestPath := []string{}

	for _, move := range state.SafeMoves() {
		nextState, err := state.Move(move)

		if err != nil {
			log.Print("Cannot move from current state")
			return []string{}, true
		}
		path, terminal := Dfs(*nextState, depth-1)

		if !terminal {
			path = append([]string{move}, path...)
			return path, false
		}
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}
	return longestPath, true
}
