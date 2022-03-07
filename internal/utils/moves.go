package utils

import (
	"battlesnake/pkg/api"
	"errors"
)

func MoveCoord(place api.Coord, move string) api.Coord {
	next := place

	// Update head
	switch move {
	case "left":
		next.X -= 1
	case "right":
		next.X += 1
	case "down":
		next.Y -= 1
	case "up":
		next.Y += 1
	}
	return next
}

func AllMoves(possible bool) map[string]bool {
	return map[string]bool{
		"up":    possible,
		"down":  possible,
		"left":  possible,
		"right": possible,
	}
}

func AvoidNeck(possibleMoves map[string]bool, snake api.Battlesnake) {
	neck := snake.Body[1]

	// Don't let your Battlesnake move back in on it's own neck
	if neck.X < snake.Head.X {
		possibleMoves["left"] = false
	} else if neck.X > snake.Head.X {
		possibleMoves["right"] = false
	} else if neck.Y < snake.Head.Y {
		possibleMoves["down"] = false
	} else if neck.Y > snake.Head.Y {
		possibleMoves["up"] = false
	}
}

func AvoidWalls(possibleMoves map[string]bool, snake api.Battlesnake, board api.Board) map[string]bool {
	// Don't hit walls.
	if snake.Head.X == 0 {
		possibleMoves["left"] = false
	}
	if snake.Head.X == board.Width-1 {
		possibleMoves["right"] = false
	}
	if snake.Head.Y == 0 {
		possibleMoves["down"] = false
	}
	if snake.Head.Y == board.Height-1 {
		possibleMoves["up"] = false
	}
	return possibleMoves
}

// Checks and avoids a list of obstacles, only tries to avoid them (according to tryRemoveMove) if soft is set
func AvoidObstacles(possibleMoves map[string]bool, snake api.Battlesnake, obstacles []api.Coord, soft bool) error {
	for _, obstacle := range obstacles {
		if snake.Head.Y == obstacle.Y {
			if snake.Head.X+1 == obstacle.X {
				RemoveMove(possibleMoves, "right", soft)
			}
			if snake.Head.X-1 == obstacle.X {
				RemoveMove(possibleMoves, "left", soft)
			}
		}
		if snake.Head.X == obstacle.X {
			if snake.Head.Y+1 == obstacle.Y {
				RemoveMove(possibleMoves, "up", soft)
			}
			if snake.Head.Y-1 == obstacle.Y {
				RemoveMove(possibleMoves, "down", soft)
			}
		}
		if snake.Head.X == obstacle.X && snake.Head.Y == obstacle.Y {
			return errors.New("Snake is in obstacle")
		}
	}
	return nil
}

func RemoveMove(possibleMoves map[string]bool, move string, soft bool) {
	if soft {
		TryRemoveMove(possibleMoves, move)
	} else {
		possibleMoves[move] = false
	}
}

func TryRemoveMove(possibleMoves map[string]bool, move string) bool {
	possibleMoves[move] = false

	// check if any are possible
	for _, possible := range possibleMoves {
		if possible {
			return true
		}
	}
	// otherwise reset
	possibleMoves[move] = true
	return false
}

// === Converters === //

func SafeMoves(possibleMoves map[string]bool) []string {
	safeMoves := []string{}

	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}
	return safeMoves
}

func PossibleMoves(safeMoves []string) map[string]bool {
	possibleMoves := AllMoves(false)

	for _, move := range safeMoves {
		possibleMoves[move] = true
	}
	return possibleMoves
}
