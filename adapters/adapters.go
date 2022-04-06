package chatadapters

type ChatPlatform interface {
	Connect() error
	GetAGUsers() []User
	IsInVoice(user User) bool
	Disconnect() error
}

type User struct {
	DiscordId string
	Name      string
}
