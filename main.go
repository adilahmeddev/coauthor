package main

import (
	"encoding/json"
	"fmt"
	"github.com/AdilahmedDev/coauthor/adapters/disc"
	coauthor "github.com/AdilahmedDev/coauthor/lib"
	"io"
	"log"
	"os"
)

func main() {
	pairSource := getSourceFromArgs(os.Args)
	filePath := "authors.json"
	var (
		pairs []string
		err   error
	)

	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

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

func loadConfig() (config disc.Config, err error) {
	set := false
	file, err := os.Open("coauthor-config.json")
	if err != nil {
		log.Println("coauthor-config.json not found, trying environment")
		config = disc.Config{
			GuildID:    os.Getenv("discord_guild_id"),
			ChannelIDA: os.Getenv("discord_channel_a"),
			BotToken:   os.Getenv("discord_bot"),
			ChannelIDB: os.Getenv("discord_channel_b"),
			MyID:       os.Getenv("my_id"),
		}
		if config.ChannelIDB == "" && config.ChannelIDA == "" {
			return disc.Config{}, fmt.Errorf("no Discord channel IDs provided")
		}
		if config.MyID == "" {
			return disc.Config{}, fmt.Errorf("MyID is not provided")
		}
		if config.GuildID == "" {
			return disc.Config{}, fmt.Errorf("GuildID is not provided")
		}
		if config.BotToken == "" {
			return disc.Config{}, fmt.Errorf("BotToken is not provided")
		}
		fmt.Println("from env")
		set = true
	} else {
		var bytes []byte
		bytes, err = io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			log.Fatal(err)
		}
		set = true
		fmt.Println("from file")
	}
	if !set {
		return disc.Config{}, fmt.Errorf("config has not been set")
	}
	return
}

func getSourceFromArgs(args []string) string {
	if len(args) > 1 {
		return args[1]

	}
	return ""
}
