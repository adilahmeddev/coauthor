package main

import (
	chatadapters "coauthor/adapters"
	"coauthor/adapters/disc"
	"fmt"
	"os"
)

func main() {
	users := []chatadapters.User{{DiscordId: os.Getenv("adil_dc_id"), GithubEmail: "adilahmeddev@gmail.com", Name: "adil"}}
	config := disc.Config{
		GuildID:    os.Getenv("discord_guild_id"),
		ChannelIDA: os.Getenv("discord_channel_a"),
		BotToken:   os.Getenv("discord_bot"),
		ChannelIDB: "",
	}
	fmt.Println(config)
	adapter := disc.NewDiscord(config, users)

	err := adapter.Connect()
	defer adapter.Disconnect()
	if err != nil {
		panic(err)
	}

	agUsers := adapter.GetAGUsers()

	fmt.Println(agUsers)

}
