package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	gamemaker "github.com/Quinn-Donnelly/discord-game-bot/game_maker"
	discord "github.com/bwmarrin/discordgo"
)

// This is the color of the embeded border on messages
var messageColor = 6413051
var NameOfSecretsFile = "secrets.json"

type SecretInfo struct {
	DiscordKey string
}

// Manages a single session of the bots ability to make game selections
var currentGameManager gamemaker.GameManager

func main() {
	secrets := readSecretsFile()

	discord, err := discord.New("Bot " + secrets.DiscordKey)

	if err != nil {
		log.Fatal(err)
		return
	}

	discord.AddHandler(messageCreated)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func readSecretsFile() SecretInfo {
	// Read in secrets
	secretsFile, err := os.Open(NameOfSecretsFile)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer secretsFile.Close()
	bytes, _ := ioutil.ReadAll(secretsFile)

	var secrets SecretInfo
	err = json.Unmarshal(bytes, &secrets)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return secrets
}

func messageCreated(session *discord.Session, messageCreateEvent *discord.MessageCreate) {
	// Don't worry about the bots own messages
	if messageCreateEvent.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(messageCreateEvent.Content, ">") {
		message := strings.Split(messageCreateEvent.Content, " ")

		if len(message) < 2 {
			session.ChannelMessageSend(messageCreateEvent.ChannelID, "Please issue a command.")
		}

		command := message[1]
		var response string
		printListWhenDone := true
		switch command {
		case "addgame":
			response = addGame(message[2:], messageCreateEvent.Author.Username)
		case "removegame":
			response = removegame(message[2:])
		case "random":
			response = pickRandom()
			printListWhenDone = false
		default:
			response = "Unkown Command."
		}

		session.ChannelMessageSend(messageCreateEvent.ChannelID, response)

		if printListWhenDone {
			printCurrentList(session, messageCreateEvent.ChannelID)
		}
	}
}

func addGame(message []string, requestedBy string) string {
	if len(message) < 1 {
		return "No game specified"
	}

	currentGameManager.AddGame(gamemaker.GameRequest{
		RequestedBy:   requestedBy,
		RequestedGame: gamemaker.Game{Name: message[0]},
	})

	return message[0] + " has been added to selections"
}

func removegame(message []string) string {
	number, err := strconv.Atoi(message[0])
	var nameOfGameRemoved string

	if err == nil {
		// Handle case where user is removing by index in list
		index := number - 1
		removedGame, err := currentGameManager.RemoveByIndex(index)

		if err != nil {
			return "Game #" + message[0] + " does not exist"
		}

		nameOfGameRemoved = removedGame.Name
	} else {
		nameOfGame := strings.ToLower(strings.Join(message, " "))

		removedGame := currentGameManager.RemoveByName(nameOfGame)

		if removedGame.Name != nameOfGame {
			return "Game with name " + nameOfGame + " does not exist"
		}

		nameOfGameRemoved = removedGame.Name
	}

	return nameOfGameRemoved + " has been removed"
}

func pickRandom() (message string) {
	game, err := currentGameManager.SelectRandomGame()
	if err != nil {
		return err.Error()
	}

	return strings.Title(game.Name) + " has been selected"
}

func printCurrentList(session *discord.Session, channelId string) {

	var fields []*discord.MessageEmbedField
	gameCount := 1

	for _, e := range currentGameManager.ListOfGames {
		fields = append(fields, &discord.MessageEmbedField{
			Name:  strconv.Itoa(gameCount) + ". " + strings.Title(e.RequestedGame.Name),
			Value: "Requested by: " + e.RequestedBy,
		})

		gameCount++
	}

	if len(fields) == 0 {
		fields = append(fields, &discord.MessageEmbedField{
			Name:  "No games have been added to list",
			Value: "To add a game use addgame command",
		})
	}

	msg := discord.MessageEmbed{
		Title:       "Games in Pool",
		Description: "These are the games that have the potentail to be selected",
		Fields:      fields,
		Color:       messageColor,
	}

	session.ChannelMessageSendEmbed(channelId, &msg)
}
