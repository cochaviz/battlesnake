package data

import (
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"container/list"
	"errors"
	"log"
)

type GameState struct {
	Turn        int
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
	*newState = state
	return newState
}

func ConvertFrom(state api.GameState) *GameState {
	newState := new(GameState)

	newState.Turn = state.Turn
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

func (state GameState) CountSpace() int32 {
	// snakes are obstacles
	obstacles := append(state.You.Body)
	for _, snake := range state.OtherSnakes {
		obstacles = append(obstacles, snake.Body...)
	}
	freeSpace := list.New()
	spaceCount := int32(0)
	freeSpace.Init()

	// start with the head
	freeSpace.PushBack(state.You.Head)

	for {
		if freeSpace.Len() == 0 {
			break
		}
		currentSpace := freeSpace.Front().Value.(api.Coord)
		possibleMoves := utils.AllMoves(true)
		mockSnake := api.Battlesnake{Head: currentSpace}

		utils.AvoidWalls(possibleMoves, mockSnake, state.Board)
		utils.AvoidObstacles(possibleMoves, mockSnake, obstacles, false)

		for _, move := range utils.SafeMoves(possibleMoves) {
			freeSpace.PushBack(utils.MoveCoord(currentSpace, move))
		}
		obstacles = append(obstacles, currentSpace)
		spaceCount++
	}
	// Because we shouldn't count the head
	spaceCount--

	return spaceCount
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
	return state.You.Health == 0 || len(state.moves) == 0
}

func (state GameState) Move(move string) (*GameState, error) {
	if state.IsTerminal() {
		return &state, errors.New("State is terminal")
	}
	nextState := state.Copy()
	nextState.You.Head = utils.MoveCoord(nextState.You.Head, move)
	nextState.You.Body = append([]api.Coord{nextState.You.Head}, nextState.You.Body...)

	if !nextState.onFood() {
		nextState.You.Body[len(nextState.You.Body)-1] = api.Coord{}
		nextState.You.Body = nextState.You.Body[:len(nextState.You.Body)-1]
		nextState.You.Health = 100
	} else {
		nextState.You.Length++
		nextState.You.Health--
	}
	nextState.updateMoves()

	return nextState, nil
}
