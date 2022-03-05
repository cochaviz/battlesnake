package data

import (
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"errors"
	"log"
)

type GameState struct {
	Board       api.Board
	OtherSnakes []api.Battlesnake
	You         api.Battlesnake
	moves       []string
}

func (state *GameState) updateMoves() {
	possibleMoves := utils.AllMoves(true)

	utils.AvoidNeck(possibleMoves, state.You)
	utils.AvoidWalls(possibleMoves, state.You, state.Board)

	// Avoid yourself
	utils.AvoidObstacles(possibleMoves, state.You, state.You.Body[1:], false)

	for _, otherSnake := range state.OtherSnakes {
		if error := utils.AvoidObstacles(possibleMoves, state.You, otherSnake.Body, false); error != nil {
			state.moves = []string{}
			log.Println("Warning: snake is in an obstacle")
			return
		}
	}
	state.moves = utils.SafeMoves(possibleMoves)
}

func (state *GameState) Init() {
	state.OtherSnakes = []api.Battlesnake{}

	for _, snake := range state.Board.Snakes {
		if snake.Head != state.You.Head {
			state.OtherSnakes = append(state.OtherSnakes, []api.Battlesnake{snake}...)
		}
	}
	state.updateMoves()
}

func (state GameState) Copy() *GameState {
	newState := new(GameState)

	newState.Board = state.Board
	newState.You = state.You
	newState.OtherSnakes = state.OtherSnakes
	newState.moves = state.moves

	return newState
}

func ConvertFrom(state api.GameState) *GameState {
	newState := new(GameState)

	newState.Board = state.Board
	newState.You = state.You
	newState.Init()

	return newState
}

func (state GameState) PossibleMoves() map[string]bool {
	return utils.PossibleMoves(state.moves)
}

func (state GameState) SafeMoves() []string {
	return state.moves
}

func (state GameState) onFood() bool {
	for _, food := range state.Board.Food {
		if food == state.You.Head {
			return true
		}
	}
	return false
}

func (state GameState) IsTerminal() bool {
	return len(state.moves) == 0
}

func (state GameState) Move(move string) (*GameState, error) {
	if state.IsTerminal() {
		return &state, errors.New("State is terminal")
	}
	nextState := state.Copy()
	newHead := api.Coord{X: nextState.You.Head.X, Y: nextState.You.Head.Y}

	// Update head
	switch move {
	case "left":
		newHead.X -= 1
	case "right":
		newHead.X += 1
	case "down":
		newHead.Y -= 1
	case "up":
		newHead.Y += 1
	}
	nextState.You.Head = newHead
	nextState.You.Body = append([]api.Coord{nextState.You.Head}, nextState.You.Body...)

	// TODO Update health
	if !nextState.onFood() {
		nextState.You.Body[len(nextState.You.Body)-1] = api.Coord{}
		nextState.You.Body = nextState.You.Body[:len(nextState.You.Body)-1]
	} else {
		nextState.You.Length++
	}
	nextState.updateMoves()

	return nextState, nil
}
