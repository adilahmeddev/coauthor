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
		os.Exit(1)
	}

	pairSource := getSourceFromArgs(os.Args)

	config, conferr := loadConfig()

	users, err := lib.GetAuthorList(authorsFilePath)
	if err != nil {
		os.Exit(1)
	}
	switch pairSource {
	case "discord":
		if conferr != nil {
			os.Exit(1)
		}
		coauthors = lib.GetPairsFromDiscord(config, users)
		if len(coauthors) == 0 {
			os.Exit(1)
		}
	case "pairs":
		coauthors, err = lib.GetPairsFromJSON(users)
	}

	output := lib.PrepareCommitMessage(string(file), coauthors)

	err = os.WriteFile(commitFilePath, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
	if len(coauthors) > 0 {
		fmt.Println("Added co-authors:", coauthors)
	} else {
		fmt.Println("No co-authors added")
	}

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
		set = true
	} else {
		var bytes []byte
		bytes, err = io.ReadAll(file)
		if err != nil {
			os.Exit(1)
		}
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			os.Exit(1)
		}
		set = true
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
