package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
)

// This is the color of the embeded border on messages
var messageColor = 6413051
var NameOfSecretsFile = "secrets.json"

type GameRequest struct {
	RequestedBy string
	Game        string
}

type SecretInfo struct {
	DiscordKey string
}

var games = list.New()

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
		switch command {
		case "addgame":
			response = addGame(message[2:], messageCreateEvent.Author.Username)
		case "removegame":
			response = removegame(message[2:])
		default:
			response = "Unkown Command."
		}

		session.ChannelMessageSend(messageCreateEvent.ChannelID, response)
		printCurrentList(session, messageCreateEvent.ChannelID)
	}
}

func addGame(message []string, requestedBy string) string {
	if len(message) < 1 {
		return "No game specified"
	}

	games.PushBack(GameRequest{
		RequestedBy: requestedBy,
		Game:        message[0],
	})
	return message[0] + " has been added to selections"
}

func removegame(message []string) string {
	return "Not Implemented"
}

func printCurrentList(session *discord.Session, channelId string) {

	var fields []*discord.MessageEmbedField
	gameCount := 1

	for e := games.Front(); e != nil; e = e.Next() {
		fields = append(fields, &discord.MessageEmbedField{
			Name:  strconv.Itoa(gameCount) + ". " + e.Value.(GameRequest).Game,
			Value: "Requested by: " + e.Value.(GameRequest).RequestedBy,
		})

		gameCount++
	}

	msg := discord.MessageEmbed{
		Title:       "Games in Pool",
		Description: "These are the games that have the potentail to be selected",
		Fields:      fields,
		Color:       messageColor,
	}

	session.ChannelMessageSendEmbed(channelId, &msg)
}
