package game_maker

import "strings"

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

// NOTE: Need to check if the idx + 1 goes out of bounds
func (g *GameManager) RemoveGame(game Game) (removedGame Game) {
	for idx, gameInList := range g.ListOfGames {
		if strings.ToLower(game.Name) == gameInList.Name {
			removedGame = gameInList
			g.ListOfGames = append(g.ListOfGames[:idx], g.ListOfGames[(idx+1):]...)
		}
	}

	return removedGame
}

// NOTE: Need to check if the idx + 1 goes out of bounds
func (g *GameManager) RemoveByName(game string) (removedGame Game) {
	for idx, gameInList := range g.ListOfGames {
		if strings.ToLower(game) == gameInList.Name {
			removedGame = gameInList
			g.ListOfGames = append(g.ListOfGames[:idx], g.ListOfGames[(idx+1):]...)
		}
	}

	return removedGame
}

// NOTE: Need to check if the idx + 1 goes out of bounds
func (g *GameManager) RemoveByIndex(game int) (removedGame Game) {
	removedGame = g.ListOfGames[game]
	g.ListOfGames = append(g.ListOfGames[:game], g.ListOfGames[game+1:]...)
	return removedGame
}
