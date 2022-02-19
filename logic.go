package main

import (
	"log"
	"math"
	"math/rand"
)

func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "cochaviz",
		Color:      "#488A16",
		Head:       "beluga",
		Tail:       "curled",
	}
}

// This function is called everytime your Battlesnake is entered into a game.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function removes moves that will get the snake further away from food if there are other options
func greedyMove(possibleMoves map[string]bool, state GameState, myHead Coord) {
	var minDistance int = math.MaxInt
	var minFood Coord

	for _, food := range state.Board.Food {
		var newDistance = manHattanDistance(myHead, food)

		if newDistance < minDistance {
			minFood = food
			minDistance = newDistance
		}
	}

	if minDistance != math.MaxInt {
		if minFood.X >= myHead.X {
			if tryRemoveMove(possibleMoves, "left") {
			}
		}
		if minFood.X <= myHead.X {
			if tryRemoveMove(possibleMoves, "right") {
			}
		}
		if minFood.Y >= myHead.Y {
			if tryRemoveMove(possibleMoves, "down") {
			}
		}
		if minFood.Y <= myHead.Y {
			if tryRemoveMove(possibleMoves, "up") {
			}
		}
	}
}

func randomMove(safeMoves []string) string {
	return safeMoves[rand.Intn(len(safeMoves))]
}

// Checks and avoids a list of obstacles, only tries to avoid them (according to tryRemoveMove) if soft is set
func checkAround(possibleMoves map[string]bool, myHead Coord, obstacle []Coord, soft bool) {
	for _, part := range obstacle {
		if myHead.Y == part.Y {
			if myHead.X+1 == part.X {
				removeMove(possibleMoves, "right", soft)
			}
			if myHead.X-1 == part.X {
				removeMove(possibleMoves, "left", soft)
			}
		}
		if myHead.X == part.X {
			if myHead.Y+1 == part.Y {
				removeMove(possibleMoves, "up", soft)
			}
			if myHead.Y-1 == part.Y {
				removeMove(possibleMoves, "down", soft)
			}
		}
	}
}

// Returns a list of moves possible without dying
func getPossibleMoves(state GameState) map[string]bool {
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	myBody := state.You.Body
	myHead := myBody[0]
	myNeck := myBody[1]
	myTail := myBody[2:]
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	// Don't let your Battlesnake move back in on it's own neck
	if myNeck.X < myHead.X {
		possibleMoves["left"] = false
	} else if myNeck.X > myHead.X {
		possibleMoves["right"] = false
	} else if myNeck.Y < myHead.Y {
		possibleMoves["down"] = false
	} else if myNeck.Y > myHead.Y {
		possibleMoves["up"] = false
	}

	// Don't hit walls.
	if myHead.X == 0 {
		possibleMoves["left"] = false
	}
	if myHead.X == boardWidth-1 {
		possibleMoves["right"] = false
	}
	if myHead.Y == 0 {
		possibleMoves["down"] = false
	}
	if myHead.Y == boardHeight-1 {
		possibleMoves["up"] = false
	}

	// Don't hit yourself.
	checkAround(possibleMoves, myHead, myTail, false)

	// Don't hit others.
	for _, others := range state.Board.Snakes {
		checkAround(possibleMoves, myHead, others.Body, false)
	}

	return possibleMoves
}

// This function is called on every turn of a game. Use the provided GameState to decide
func move(state GameState) BattlesnakeMoveResponse {
	possibleMoves := getPossibleMoves(state)
	myHead := state.You.Body[0]

	// Only find food if we need it
	if state.You.Health < 25 {
		greedyMove(possibleMoves, state, myHead)
	} else {
		// Otherwise try to avoid it
		checkAround(possibleMoves, myHead, state.Board.Food, true)
	}

	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving down\n", state.Game.ID, state.Turn)
	} else {
		nextMove = randomMove(safeMoves)
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
