package internal

import (
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"log"
	"math"
	"math/rand"
)

// Attributes a value according to the state of the game [math.MinInt, math.MaxInt]
func heuristic(state api.GameState) int {
	if len(getPossibleMoves(state)) == 0 {
		return math.MinInt
	}
	return math.MaxInt
}

// This function removes moves that will get the snake further away from food if there are other options
func greedyMove(possibleMoves map[string]bool, state api.GameState, myHead api.Coord) {
	var minDistance int = math.MaxInt
	var minFood api.Coord

	for _, food := range state.Board.Food {
		var newDistance = utils.ManHattanDistance(myHead, food)

		if newDistance < minDistance {
			minFood = food
			minDistance = newDistance
		}
	}

	if minDistance != math.MaxInt {
		if minFood.X >= myHead.X {
			if utils.TryRemoveMove(possibleMoves, "left") {
			}
		}
		if minFood.X <= myHead.X {
			if utils.TryRemoveMove(possibleMoves, "right") {
			}
		}
		if minFood.Y >= myHead.Y {
			if utils.TryRemoveMove(possibleMoves, "down") {
			}
		}
		if minFood.Y <= myHead.Y {
			if utils.TryRemoveMove(possibleMoves, "up") {
			}
		}
	}
}

// Returns a list of moves possible without dying
func getPossibleMoves(state api.GameState) map[string]bool {
	myBody := state.You.Body
	myHead := myBody[0]
	myNeck := myBody[1]
	myTail := myBody[2:]
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	possibleMoves := utils.AllMoves()

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
	utils.PossibleMovesWithinBounds(possibleMoves, myHead, boardWidth, boardHeight)
	log.Println("Possible moves within bounds: ", possibleMoves)

	// Don't hit yourself.
	utils.CheckAround(possibleMoves, myHead, myTail, false)

	// Don't hit others.
	for _, others := range state.Board.Snakes {
		utils.CheckAround(possibleMoves, myHead, others.Body, false)
	}

	return possibleMoves
}

func Think(state api.GameState) string {
	possibleMoves := getPossibleMoves(state)
	myHead := state.You.Body[0]

	// Only find food if we need it
	if state.You.Health < 25 {
		greedyMove(possibleMoves, state, myHead)
	} else {
		// Otherwise try to avoid it
		utils.CheckAround(possibleMoves, myHead, state.Board.Food, true)
	}

	var nextMove string
	safeMoves := utils.GetSafeMoves(possibleMoves)

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving down\n", state.Game.ID, state.Turn)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))] // Pick random move from the safe moves
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return nextMove
}
