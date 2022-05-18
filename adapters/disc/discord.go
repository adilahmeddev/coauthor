package disc

import (
	chatadapters "github.com/AdilahmedDev/coauthor/adapters"
	"github.com/alecthomas/repr"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Config struct {
	GuildID    string
	ChannelIDA string
	ChannelIDB string
	BotToken   string
	MyID       string
}

type Discord struct {
	Session *discordgo.Session
	config  Config
	authors []chatadapters.User
}

func NewDiscord(config Config, users []chatadapters.User) *Discord {
	return &Discord{config: config, authors: users}
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
func isAuthor(slice []chatadapters.User, id string) bool {
	for _, t := range slice {
		if t.DiscordId == id {
			return true
		}
	}
	return false
}
func (d *Discord) GetAGUsers() []chatadapters.User {

	returnUsers := []chatadapters.User{}
	me, err := d.Session.GuildMember(d.config.GuildID, d.config.MyID)
	if err != nil {
		log.Fatal("cant get your user")
	}
	myState, err := d.Session.State.VoiceState(d.config.GuildID, me.User.ID)
	if err != nil {
		log.Fatal("Unable to get your voice session. Are you in a voice chat in the right server?")
	}
	for _, user := range d.authors {
		if user.DiscordId == "" {
			continue
		}
		member, err := d.Session.GuildMember(d.config.GuildID, user.DiscordId)
		if err != nil {
			log.Fatal("could not get ", repr.String(user), "member info ", err)
		}
		if member != nil {
			state, err := d.Session.State.VoiceState(d.config.GuildID, member.User.ID)
			if err != nil {
				log.Fatal("could not get users state", err)
			}
			if state != nil && myState != nil {

				if (state.ChannelID == d.config.ChannelIDA || state.ChannelID == d.config.ChannelIDB) && user.DiscordId != d.config.MyID && isAuthor(d.authors, user.DiscordId) {
					returnUsers = append(returnUsers, user)
				}
			}
		}
	}
	return returnUsers
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
