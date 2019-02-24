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
	ListOfGames []GameRequest
}

func (g *GameManager) AddGame(game GameRequest) {
	game.RequestedGame.Name = strings.ToLower(game.RequestedGame.Name)
	g.ListOfGames = append(g.ListOfGames, game)
}

func (g *GameManager) RemoveGame(game Game) (removedGame Game) {
	for idx, gameInList := range g.ListOfGames {
		if strings.ToLower(game.Name) == gameInList.RequestedGame.Name {
			removedGame = gameInList.RequestedGame
			g.ListOfGames = append(g.ListOfGames[:idx], g.ListOfGames[(idx+1):]...)
		}
	}

	return removedGame
}

func (g *GameManager) RemoveByName(game string) (removedGame Game) {
	for idx, gameInList := range g.ListOfGames {
		if strings.ToLower(game) == gameInList.RequestedGame.Name {
			removedGame = gameInList.RequestedGame
			g.ListOfGames = append(g.ListOfGames[:idx], g.ListOfGames[(idx+1):]...)
		}
	}

	return removedGame
}

// TODO: Update a test to check for the error condition that was added
func (g *GameManager) RemoveByIndex(game int) (removedGame Game, err error) {
	if game > len(g.ListOfGames) || game < 0 {
		return Game{}, errors.New("Index out of bounds")
	}

	removedGame = g.ListOfGames[game].RequestedGame
	g.ListOfGames = append(g.ListOfGames[:game], g.ListOfGames[game+1:]...)
	return removedGame, nil
}

func (g *GameManager) SelectRandomGame() (selectedGame Game, err error) {
	if len(g.ListOfGames) < 1 {
		return Game{}, errors.New("No games to choose from.")
	}

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randomGenerator := rand.New(source)

	index := randomGenerator.Intn(len(g.ListOfGames))
	selectedGame = g.ListOfGames[index].RequestedGame
	return selectedGame, nil
}
