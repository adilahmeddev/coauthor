package disc

import (
	"fmt"
	chatadapters "github.com/AdilahmedDev/coauthor/adapters"
	"github.com/bwmarrin/discordgo"
)

type Config struct {
	GuildID    string
	ChannelIDA string
	ChannelIDB string
	BotToken   string
}

type Discord struct {
	Session *discordgo.Session
	config  Config
	users   []chatadapters.User
}

func NewDiscord(config Config, users []chatadapters.User) *Discord {
	return &Discord{config: config, users: users}
}

func (d *Discord) Connect() error {
	d.Session, _ = discordgo.New("Bot " + d.config.BotToken)

	d.Session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMembers
	err := d.Session.Open()

	if err != nil {
		return err
	}
	return nil
}

func (d *Discord) Disconnect() error {
	return d.Session.Close()
}

func (d *Discord) GetAGUsers() []chatadapters.User {
	users := []chatadapters.User{}
	for _, user := range d.users {
		member, _ := d.Session.GuildMember(d.config.GuildID, user.DiscordId)

		state, _ := d.Session.State.VoiceState(d.config.GuildID, member.User.ID)

		if state != nil {
			if state.ChannelID == d.config.ChannelIDA || state.ChannelID == d.config.ChannelIDB {
				users = append(users, user)
			}
		}

		fmt.Printf("%+v\n", member.User.Email)
	}
	return users
}

func (d *Discord) IsInVoice(user chatadapters.User) bool {
	member, err := d.Session.GuildMember(d.config.GuildID, user.DiscordId)
	if err != nil {
		panic(err)
	}

	state, _ := d.Session.State.VoiceState(d.config.GuildID, member.User.ID)
	if state != nil {

	}

	if state != nil {
		return true
	}
	return false
}
