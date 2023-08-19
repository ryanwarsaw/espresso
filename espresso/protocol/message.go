package protocol

import (
	"errors"
	"strings"
)

type Message struct {
	Source     Source
	Command    string
	Parameters []string
}

type Source struct {
	ServerName string
	Nick       string
	User       string
	Host       string
}

func ParseSource(source string) (Source, error) {
	splitOnHost := strings.Split(source, "@")
	if len(splitOnHost) == 2 {
		splitOnUser := strings.Split(splitOnHost[0], "!")
		if len(splitOnUser) != 2 {
			return Source{}, errors.New("invalid source format")
		}

		return Source{
			ServerName: "",
			Host:       splitOnHost[1],
			Nick:       splitOnUser[0],
			User:       splitOnUser[1],
		}, nil
	}

	return Source{
		ServerName: splitOnHost[0],
		Host:       "",
		Nick:       "",
		User:       "",
	}, nil
}

// Parses an IRCv3 protocol message for further processing by individual handlers
func ParseMessage(message string) (Message, error) {
	if len(message) > 0 {
		var source Source
		var command string
		var parameters []string

		hasSourcePrefix := message[0] == ':'
		args := strings.Fields(message)
		parseOffset := 0

		if hasSourcePrefix {
			result, err := ParseSource(args[0][1:])
			if err != nil {
				return Message{}, err
			}
			source = result
			parseOffset = 1
		}

		command = args[parseOffset]
		parameters = args[parseOffset+1:]

		return Message{
			Source:     source,
			Command:    command,
			Parameters: parameters,
		}, nil
	}
	return Message{}, errors.New("empty message")
}
