package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomAttributes map[string]interface{}

type snakeData struct {
	Body  []Coord `json:"body"`
	Head  Coord   `json:"head"`
	Color string  `json:"color"`
}

func snakeDataFrom(snake Battlesnake) snakeData {
	// TODO Actual color
	return snakeData{snake.Body, snake.Head, "#000000"}
}

func base64Json(v interface{}) (string, error) {
	jsonEncoded, err := json.Marshal(v)

	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(jsonEncoded), nil
}

func GetCustomAttributesFromGameState(state GameState) (CustomAttributes, error) {
	customAttributes := CustomAttributes{}

	// -- Game properties --
	customAttributes["snakeGameId"] = state.Game.ID
	customAttributes["snakeRules"] = state.Game.Ruleset.Name
	customAttributes["snakeTurn"] = state.Turn

	// -- Board properties --
	customAttributes["snakeBoardHeight"] = state.Board.Height
	customAttributes["snakeBoardWidth"] = state.Board.Width

	// Encode to json, and then base64
	snakeBoardFood, err := base64Json(state.Board.Food)
	if err != nil {
		return CustomAttributes{}, err
	}
	customAttributes["snakeBoardFood"] = snakeBoardFood

	// Encode to json, and then base64
	snakeBoardHazards, err := base64Json(state.Board.Hazards)
	if err != nil {
		return CustomAttributes{}, err
	}
	customAttributes["snakeBoardHazards"] = snakeBoardHazards

	// -- You properties --
	customAttributes["snakeName"] = state.You.Name
	customAttributes["snakeId"] = state.You.ID
	customAttributes["snakeHealth"] = state.You.Health
	customAttributes["snakeLength"] = state.You.Length

	// Encode to json, and then base64
	encodedSnakeData, err := base64Json(snakeDataFrom(state.You))
	if err != nil {
		return CustomAttributes{}, nil
	}
	customAttributes["snakeData"] = encodedSnakeData

	// -- Opponent properties

	for index, snake := range state.Board.Snakes {
		if snake.ID != state.You.ID {
			customAttributes[fmt.Sprintf("snakeOpponent_%d_snakeId", index)] = snake.ID
			customAttributes[fmt.Sprintf("snakeOpponent_%d_snakeHealth", index)] = snake.Health
			customAttributes[fmt.Sprintf("snakeOpponent_%d_snakeLength", index)] = snake.Length
			customAttributes[fmt.Sprintf("snakeOpponent_%d_snakeName", index)] = snake.Name

			// Encode to json, and then base64
			encodedSnakeData, err := base64Json(snakeDataFrom(snake))
			if err != nil {
				return CustomAttributes{}, err
			}
			customAttributes[fmt.Sprintf("snakeOpponent_%d_snakeData", index)] = encodedSnakeData
		}
	}
	return customAttributes, nil
}

func GetCustomAttributes(r *http.Request) (CustomAttributes, error) {
	state := GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)

	if err != nil {
		return CustomAttributes{}, err
	}
	return GetCustomAttributesFromGameState(state)
}
