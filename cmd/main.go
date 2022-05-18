package main

import (
	"fmt"
	"github.com/AdilahmedDev/coauthor"
	"github.com/AdilahmedDev/coauthor/adapters/disc"
	"log"
	"os"
)

func main() {
	pairSource := getSourceFromArgs(os.Args)
	filePath := "authors.json"
	config := disc.Config{
		GuildID:    os.Getenv("discord_guild_id"),
		ChannelIDA: os.Getenv("discord_channel_a"),
		BotToken:   os.Getenv("discord_bot"),
		ChannelIDB: "",
		MyID:       os.Getenv("my_id"),
	}
	var (
		pairs []string
		err   error
	)
	users, err := coauthor.GetAuthorList(filePath)
	if err != nil {
		log.Fatal(err)
	}
	switch pairSource {
	case "discord":
		pairs = coauthor.GetPairsFromDiscord(config, users)
	case "pairs":
		pairs, err = coauthor.GetPairsFromJSON(users)
	}

	fmt.Println(pairs)
}

func getSourceFromArgs(args []string) string {
	if len(args) > 1 {
		return args[1]

	}
	return ""
}
