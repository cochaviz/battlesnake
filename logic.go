package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

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
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

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
			log.Printf("Current possibleMoves: %v", possibleMoves)
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

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
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

	// Find food.
	if state.You.Health < 25 {
		greedyMove(possibleMoves, state, myHead)
	} else {
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
