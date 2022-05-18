package main

import (
	"encoding/json"
	"flag"
	"fmt"
	chatadapters "github.com/AdilahmedDev/coauthor/adapters"
	"github.com/AdilahmedDev/coauthor/adapters/disc"
	"github.com/AdilahmedDev/coauthor/lib"
	"io"
	"log"
	"os"
)

var (
	authorsFilePath string
	commitFilePath  string
	pairsFilePath   string
)

func init() {
	flag.StringVar(&authorsFilePath, "authorsFile", "authors.json", "names & emails of teammates")
	flag.StringVar(&commitFilePath, "commitFile", ".git/COMMIT_EDITMSG", "path to commit message file")
	flag.StringVar(&pairsFilePath, "pairsFile", "pairs.json", "path to pairs file")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	var (
		coauthors chatadapters.Users
		err       error
	)

	file, err := os.ReadFile(commitFilePath)
	if err != nil {
		log.Fatal(err)
	}

	pairSource := getSourceFromArgs(os.Args)

	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	users, err := lib.GetAuthorList(authorsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	switch pairSource {
	case "discord":
		coauthors = lib.GetPairsFromDiscord(config, users)
	case "pairs":
		coauthors, err = lib.GetPairsFromJSON(users)
	}

	output := lib.PrepareCommitMessage(string(file), coauthors)

	err = os.WriteFile(commitFilePath, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added co-authors:", coauthors)
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
