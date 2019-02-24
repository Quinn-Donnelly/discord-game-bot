package game_maker

import (
	"errors"
	"testing"
)

func TestAddGame(t *testing.T) {
	gameMaker := GameManager{
		ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}, RequestedBy: "quinn"}},
	}

	gameMaker.AddGame(GameRequest{
		RequestedGame: Game{Name: "dota"},
	})

	found := false
	for _, game := range gameMaker.ListOfGames {
		if game.RequestedGame.Name == "dota" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected to find dota in list but only found {}", gameMaker.ListOfGames)
	}
}

func TestAddGameDifferentCase(t *testing.T) {
	gameMaker := GameManager{
		ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}, RequestedBy: "quinn"}},
	}

	gameMaker.AddGame(GameRequest{
		RequestedGame: Game{Name: "Dota"},
	})

	found := false
	for _, game := range gameMaker.ListOfGames {
		if game.RequestedGame.Name == "dota" {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("Expected to find dota in list but only found %v", gameMaker.ListOfGames)
	}
}

func TestRemoveGameByIndex(t *testing.T) {
	testData := []struct {
		incoming       GameManager
		remove         int
		expected       GameManager
		expectedReturn Game
	}{
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}, RequestedBy: "quinn"}, GameRequest{RequestedGame: Game{Name: "dota"}, RequestedBy: "justin"}},
			},
			remove: 0,
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "dota"}, RequestedBy: "justin"}},
			},
			expectedReturn: Game{Name: "league"},
		},
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			remove: 1,
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}},
			},
			expectedReturn: Game{Name: "dota"},
		},
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}, GameRequest{RequestedGame: Game{Name: "civ"}}},
			},
			remove: 1,
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "civ"}}},
			},
			expectedReturn: Game{Name: "dota"},
		},
	}

	for _, testCase := range testData {
		removed := testCase.incoming.RemoveByIndex(testCase.remove)

		for i, _ := range testCase.incoming.ListOfGames {
			if testCase.incoming.ListOfGames[i] != testCase.expected.ListOfGames[i] {
				t.Fatalf("Expected list of games to be %+v but recieved %+v", testCase.expected, testCase.incoming)
			}
		}

		if removed != testCase.expectedReturn {
			t.Fatalf("Expected return of %v but received %v", testCase.expectedReturn, removed)
		}
	}
}

func TestRemoveByName(t *testing.T) {
	testData := []struct {
		incoming       GameManager
		remove         string
		expected       GameManager
		expectedReturn Game
	}{
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			remove: "League",
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			expectedReturn: Game{Name: "league"},
		},
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			remove: "dota",
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}},
			},
			expectedReturn: Game{Name: "dota"},
		},
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}, GameRequest{RequestedGame: Game{Name: "civ"}}},
			},
			remove: "Dota",
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "civ"}}},
			},
			expectedReturn: Game{Name: "dota"},
		},
	}

	for _, testCase := range testData {
		removed := testCase.incoming.RemoveByName(testCase.remove)

		for i, _ := range testCase.incoming.ListOfGames {
			if testCase.incoming.ListOfGames[i] != testCase.expected.ListOfGames[i] {
				t.Fatalf("Expected list of games to be %+v but recieved %+v", testCase.expected, testCase.incoming)
			}
		}

		if removed != testCase.expectedReturn {
			t.Fatalf("Expected return of %v but received %v", testCase.expectedReturn, removed)
		}
	}
}

func TestRemoveGame(t *testing.T) {
	testData := []struct {
		incoming       GameManager
		remove         Game
		expected       GameManager
		expectedReturn Game
	}{
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			remove: Game{Name: "League"},
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			expectedReturn: Game{Name: "league"},
		},
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}},
			},
			remove: Game{Name: "dota"},
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}},
			},
			expectedReturn: Game{Name: "dota"},
		},
		{
			incoming: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "dota"}}, GameRequest{RequestedGame: Game{Name: "civ"}}},
			},
			remove: Game{Name: "Dota"},
			expected: GameManager{
				ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}, GameRequest{RequestedGame: Game{Name: "civ"}}},
			},
			expectedReturn: Game{Name: "dota"},
		},
	}

	for _, testCase := range testData {
		removed := testCase.incoming.RemoveGame(testCase.remove)

		for i, _ := range testCase.incoming.ListOfGames {
			if testCase.incoming.ListOfGames[i] != testCase.expected.ListOfGames[i] {
				t.Fatalf("Expected list of games to be %+v but recieved %+v", testCase.expected, testCase.incoming)
			}
		}

		if removed != testCase.expectedReturn {
			t.Fatalf("Expected return of %v but received %v", testCase.expectedReturn, removed)
		}
	}
}

func TestSelectingGame(t *testing.T) {
	testData := []struct {
		Input       GameManager
		Expected    Game
		ExpectedErr error
	}{
		{
			Input:       GameManager{ListOfGames: []GameRequest{GameRequest{RequestedGame: Game{Name: "league"}}}},
			Expected:    Game{Name: "league"},
			ExpectedErr: nil,
		},
		{
			Input:       GameManager{ListOfGames: []GameRequest{}},
			Expected:    Game{},
			ExpectedErr: errors.New("No games to choose from."),
		},
	}

	for _, test := range testData {
		randomGame, err := test.Input.SelectRandomGame()

		if test.ExpectedErr != nil && err == nil {
			t.Fatalf("Expected error to be %v but got %v", test.ExpectedErr, err)
		}

		if randomGame != test.Expected {
			t.Fatalf("Expected RandomGame to be %+v but got %+v", test.Expected, randomGame)
		}
	}
}
