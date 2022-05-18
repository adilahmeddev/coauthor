package coauthor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	chatadapters "github.com/AdilahmedDev/coauthor/adapters"
	"github.com/AdilahmedDev/coauthor/adapters/disc"
	"io"
	"os"
	"strings"
)

func GetPairsFromDiscord(config disc.Config, users []chatadapters.User) []string {
	adapter := disc.NewDiscord(config, users)

	err := adapter.Connect()
	defer adapter.Disconnect()
	if err != nil {
		panic(err)
	}

	agUsers := adapter.GetAGUsers()
	userStrings := []string{}
	for _, user := range agUsers {
		userStrings = append(userStrings, user.Name)
	}
	return userStrings
}

func GetAuthorList(filePath string) (users []chatadapters.User, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, fmt.Errorf("please add an authors.json to your projects root directory")
		} else {
			return nil, err
		}
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}

	for i, user := range users {
		id, err := base64.StdEncoding.DecodeString(user.DiscordId)
		if err != nil {
			panic(err)
		}
		users[i].DiscordId = string(id)
	}
	return users, nil
}
func GetPairsFromJSON(authors []chatadapters.User) (coauthors []string, err error) {
	pairFile, err := os.ReadFile("pairs.json")
	if err != nil {
		return []string{}, err
	}

	var pairs []string
	err = json.NewDecoder(bytes.NewReader(pairFile)).Decode(&pairs)
	if err != nil {
		return []string{}, err
	}
	for _, pair := range pairs {
		for _, author := range authors {
			if pair == author.Name {
				coauthors = append(coauthors, pair)
			}
		}
	}
	return coauthors, nil
}
