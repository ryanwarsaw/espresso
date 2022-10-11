package protocol

import (
	"errors"
	"strings"
)

type Message struct {
	Tags       map[string]string
	Source     string
	Command    string
	Parameters []string
}

// Parses an IRCv3 protocol message into it's core components
// and constructs a message object that can be further processed
func ParseMessage(message string) (Message, error) {
	if len(message) > 0 {
		tags := make(map[string]string)
		var command string
		var source string
		var parameters []string

		args := strings.Fields(message)
		for i := 0; i < len(args); i++ {
			if args[i] == "CAP" || args[i] == "PRIVMSG" {
				command = args[i]
				parameters = args[i+1:]
				break
			}
			// Handle tags
			if args[i][0] == '@' {
				rawTags := strings.Split(args[i][1:len(args[i])], ";")
				for _, tag := range rawTags {
					keyValuePair := strings.Split(tag, "=")
					if len(keyValuePair) == 2 {
						tags[keyValuePair[0]] = keyValuePair[1]
					} else {
						return Message{}, errors.New("invalid tag format")
					}
				}
				continue
			}
			// Handle source
			if args[i][0] == ':' {
				source = args[i][1:len(args[i])]
				continue
			}
		}
		return Message{
			Tags:       tags,
			Source:     source,
			Command:    command,
			Parameters: parameters,
		}, nil
	}
	return Message{}, errors.New("empty message")
}
