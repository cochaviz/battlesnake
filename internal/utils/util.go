package utils

import (
	"battlesnake/pkg/api"
	"container/list"
	"fmt"
)

func AllMoves() map[string]bool {
	return map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
}

func IntAbs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func ManHattanDistance(a api.Coord, b api.Coord) int {
	return IntAbs(a.X-b.X) + IntAbs(a.Y-b.Y)
}

func RemoveMove(possibleMoves map[string]bool, move string, soft bool) {
	if soft {
		TryRemoveMove(possibleMoves, move)
	} else {
		possibleMoves[move] = false
	}
}

func TryRemoveMove(possibleMoves map[string]bool, move string) bool {
	switch move {
	case "left":
		if possibleMoves["right"] || possibleMoves["up"] || possibleMoves["down"] {
			possibleMoves["left"] = false
			return true
		}
		return false
	case "right":
		if possibleMoves["left"] || possibleMoves["up"] || possibleMoves["down"] {
			possibleMoves["right"] = false
			return true
		}
		return false
	case "up":
		if possibleMoves["right"] || possibleMoves["left"] || possibleMoves["down"] {
			possibleMoves["up"] = false
			return true

		}
		return false
	case "down":
		if possibleMoves["right"] || possibleMoves["up"] || possibleMoves["left"] {
			possibleMoves["down"] = false
			return true
		}
		return false
	}
	return false
}

func PossibleMovesWithinBounds(possibleMoves map[string]bool, initial api.Coord, boardWidth int, boardHeight int) {
	// Don't hit walls.
	if initial.X == 0 {
		possibleMoves["left"] = false
	}
	if initial.X == boardWidth-1 {
		possibleMoves["right"] = false
	}
	if initial.Y == 0 {
		possibleMoves["down"] = false
	}
	if initial.Y == boardHeight-1 {
		possibleMoves["up"] = false
	}
}

func GetSafeMoves(possibleMoves map[string]bool) []string {
	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}
	return safeMoves
}

// Checks and avoids a list of obstacles, only tries to avoid them (according to tryRemoveMove) if soft is set
func CheckAround(possibleMoves map[string]bool, myHead api.Coord, obstacle []api.Coord, soft bool) {
	for _, part := range obstacle {
		if myHead.Y == part.Y {
			if myHead.X+1 == part.X {
				RemoveMove(possibleMoves, "right", soft)
			}
			if myHead.X-1 == part.X {
				RemoveMove(possibleMoves, "left", soft)
			}
		}
		if myHead.X == part.X {
			if myHead.Y+1 == part.Y {
				RemoveMove(possibleMoves, "up", soft)
			}
			if myHead.Y-1 == part.Y {
				RemoveMove(possibleMoves, "down", soft)
			}
		}
		// if myHead.X == part.X && myHead.Y == part.Y {
		// 	possibleMoves["left"] = false
		// 	possibleMoves["right"] = false
		// 	possibleMoves["down"] = false
		// 	possibleMoves["up"] = false
		// }
	}
}

func CountSpace(initial api.Coord, obstacles []api.Coord, boardHeight int, boardWidth int) int {
	queue := list.New()
	var totalSpace = 1
	queue.PushBack(initial)

	for i := 0; i < 100; i++ {
		if queue.Len() == 0 {
			break
		}
		space := queue.Front().Value.(api.Coord)
		queue.Remove(queue.Front())

		fmt.Printf("Found space at %d, %d\n", space.X, space.Y)
		totalSpace += 1

		possibleMoves := AllMoves()
		PossibleMovesWithinBounds(possibleMoves, space, boardWidth, boardHeight)
		CheckAround(possibleMoves, space, obstacles, false)
		obstacles = append(obstacles, space)

		for _, move := range GetSafeMoves(possibleMoves) {
			newMove := api.Coord{X: space.X, Y: space.Y}
			switch move {
			case "up":
				newMove.Y++
			case "down":
				newMove.Y--
			case "left":
				newMove.X--
			case "right":
				newMove.X++
			}
			fmt.Printf("Enqueueing %d, %d\n", newMove.X, newMove.Y)
			queue.PushBack(newMove)
		}
	}
	return totalSpace
}
