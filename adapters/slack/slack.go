package slack

import (
	chatadapters "github.com/AdilahmedDev/coauthor/adapters"
)

type Slack struct {
}

//
//func (s Slack) Connect() error {
//	api := slack.New("YOUR_TOKEN_HERE")
//	// If you set debugging, it will log all requests to the console
//	// Useful when encountering issues
//	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
//	groups, err := api.GetUserGroups(false)
//	if err != nil {
//		fmt.Printf("%s\n", err)
//		return
//	}
//	for _, group := range groups {
//		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
//	}
//}

func (s Slack) GetAGUsers() []chatadapters.User {
	//TODO implement me
	panic("implement me")
}

func (s Slack) IsInVoice(user chatadapters.User) bool {
	//TODO implement me
	panic("implement me")
}

func (s Slack) Disconnect() error {
	//TODO implement me
	panic("implement me")
}
