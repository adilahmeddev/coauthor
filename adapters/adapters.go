package chatadapters

import "fmt"

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

func (c User) String() string {
	return fmt.Sprintf("%s <%s>", c.Name, c.Email)
}

type Users []User

func (authors Users) Get(name string) (User, error) {
	for _, author := range authors {
		if author.Name == name {
			return author, nil
		}
	}

	return User{}, fmt.Errorf("author %s not present in the authors file", name)
}
