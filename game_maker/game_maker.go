package game_maker

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	Name string
}

type GameRequest struct {
	RequestedBy   string
	RequestedGame Game
}

type GameManager struct {
	ListOfGames []Game
}

func (g *GameManager) AddGame(game Game) {
	game.Name = strings.ToLower(game.Name)
	g.ListOfGames = append(g.ListOfGames, game)
}

func (g *GameManager) RemoveGame(game Game) (removedGame Game) {
	for idx, gameInList := range g.ListOfGames {
		if strings.ToLower(game.Name) == gameInList.Name {
			removedGame = gameInList
			g.ListOfGames = append(g.ListOfGames[:idx], g.ListOfGames[(idx+1):]...)
		}
	}

	return removedGame
}

func (g *GameManager) RemoveByName(game string) (removedGame Game) {
	for idx, gameInList := range g.ListOfGames {
		if strings.ToLower(game) == gameInList.Name {
			removedGame = gameInList
			g.ListOfGames = append(g.ListOfGames[:idx], g.ListOfGames[(idx+1):]...)
		}
	}

	return removedGame
}

func (g *GameManager) RemoveByIndex(game int) (removedGame Game) {
	removedGame = g.ListOfGames[game]
	g.ListOfGames = append(g.ListOfGames[:game], g.ListOfGames[game+1:]...)
	return removedGame
}

func (g *GameManager) SelectRandomGame() (selectedGame Game, err error) {
	if len(g.ListOfGames) < 1 {
		return Game{}, errors.New("No games to choose from.")
	}

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randomGenerator := rand.New(source)

	index := randomGenerator.Intn(len(g.ListOfGames))
	selectedGame = g.ListOfGames[index]
	return selectedGame, nil
}
