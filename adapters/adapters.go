package chatadapters

type ChatPlatform interface {
	Connect() error
	GetAGUsers() []User
	IsInVoice(user User) bool
	Disconnect() error
}

type User struct {
	Name      string
	DiscordId string
	Email     string
}
