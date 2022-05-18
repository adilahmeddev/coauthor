package lib

import (
	"fmt"
	chatadapters "github.com/AdilahmedDev/coauthor/adapters"
	"strings"
)

const COMMIT_SEPARATOR = "\n# ------------------------ >8 ------------------------"

func PrepareCommitMessage(input string, coAuthors []chatadapters.User) string {
	if len(coAuthors) == 0 {
		return input
	}

	sections := strings.SplitN(input, COMMIT_SEPARATOR, 2)
	message := sections[0] + "\n"

	for _, author := range coAuthors {
		message += fmt.Sprintf("\nCo-authored-by: %s", author.String())
	}

	if len(sections) > 1 {
		metadataSection := sections[1]
		message = message + COMMIT_SEPARATOR + metadataSection
	}

	return message
}
