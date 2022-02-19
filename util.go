package main

func intAbs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func manHattanDistance(a Coord, b Coord) int {
	return intAbs(a.X-b.X) + intAbs(a.Y-b.Y)
}

func removeMove(possibleMoves map[string]bool, move string, soft bool) {
	if soft {
		tryRemoveMove(possibleMoves, move)
	} else {
		possibleMoves[move] = false
	}
}

func tryRemoveMove(possibleMoves map[string]bool, move string) bool {
	if move == "left" && (possibleMoves["right"] || possibleMoves["up"] || possibleMoves["down"]) {
		possibleMoves["left"] = false
		return true
	}
	if move == "right" && (possibleMoves["left"] || possibleMoves["up"] || possibleMoves["down"]) {
		possibleMoves["right"] = false
		return true
	}
	if move == "up" && (possibleMoves["right"] || possibleMoves["left"] || possibleMoves["down"]) {
		possibleMoves["up"] = false
		return true
	}
	if move == "down" && (possibleMoves["right"] || possibleMoves["up"] || possibleMoves["left"]) {
		possibleMoves["down"] = false
		return true
	}
	return false
}
