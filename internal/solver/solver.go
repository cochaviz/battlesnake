package solver

import (
	"battlesnake/internal/data"
	"battlesnake/internal/utils"
	"math"
)

func terminalPenalty(state data.GameState) int32 {
	if state.IsTerminal() {
		return math.MaxInt32/2 - int32(state.Turn)
	}
	return -int32(state.Turn)
}

func dislikeWalls(state data.GameState, penalty int32) int32 {
	if state.You.Head.X == 0 ||
		state.You.Head.X == state.Board.Width-1 ||
		state.You.Head.Y == 0 ||
		state.You.Head.Y == state.Board.Height-1 {
		return penalty
	}
	return 0
}

func minDistanceToTail(state data.GameState) int32 {
	return int32(utils.ManHattanDistance(state.You.Head, state.You.Body[len(state.You.Body)-1]))
}

func maxDistanceFromOthers(state data.GameState) int32 {
	totalDistance := int32(0)

	for _, snake := range state.OtherSnakes {
		totalDistance += int32(utils.ManHattanDistance(state.You.Head, snake.Head))
	}
	return totalDistance
}

func minHealth(state data.GameState) int32 {
	return state.You.Health
}

func minLength(state data.GameState) int32 {
	return state.You.Length
}

func maxLength(state data.GameState) int32 {
	return -state.You.Length
}

func cost(state data.GameState) int32 {
	return minDistanceToTail(state) +
		minLength(state) +
		maxDistanceFromOthers(state) +
		dislikeWalls(state, 5)
}

// Will try to find the longest path which is not terminal
func Dfs(state data.GameState, depth int) ([]string, int32) {
	bestPath := []string{}

	if depth == 0 || state.IsTerminal() {
		return bestPath, terminalPenalty(state)
	}
	lowestCost := int32(math.MaxInt32)
	currentCost := cost(state)

	for _, move := range state.SafeMoves() {
		nextState, _ := state.Move(move)
		path, pathCost := Dfs(*nextState, depth-1)
		nextCost := currentCost + pathCost

		if nextCost < lowestCost {
			bestPath = append([]string{move}, path...)
			lowestCost = nextCost
		}
	}
	return bestPath, lowestCost
}
