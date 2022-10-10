package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"ryanwarsaw.com/protocol"
	"strings"
)

type ConnectionOptions struct {
	Address  *string
	Port     *int
	Username *string
}

func ConfigureAndParseFlags() ConnectionOptions {
	addressPtr := flag.String("address", "localhost", "IRC server address")
	portPtr := flag.Int("port", 6667, "IRC server port")
	usernamePtr := flag.String("username", "ryan", "Personal identifier")
	flag.Parse()
	return ConnectionOptions{
		Address:  addressPtr,
		Port:     portPtr,
		Username: usernamePtr,
	}
}

type Message struct {
	Tags       map[string]string
	Source     string
	Command    string
	Parameters []string
}

// TODO: Clean this up
func parseMessage(message string) (Message, error) {
	if len(message) > 0 {
		tags := make(map[string]string)
		var command string
		var source string
		var parameters []string

		args := strings.Fields(message)
		for i := 0; i < len(args); i++ {
			if args[i] == "CAP" {
				command = "CAP"
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
						return Message{}, errors.New("invalid key/value pair in message tags")
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

func main() {
	options := ConfigureAndParseFlags()

	fmt.Println("Connecting to server with address:", *options.Address)
	fmt.Println("Using port:", *options.Port)
	fmt.Println("Using username:", *options.Username)

	hostAddress := fmt.Sprintf("%s:%d", *options.Address, *options.Port)

	connection, err := net.Dial("tcp4", hostAddress)
	if err != nil {
		log.Fatal("Failed to establish connection to host server\n", err)
	}

	connection.Write([]byte("CAP LS 302" + "\r\n"))

	for {
		message, _, err := bufio.NewReader(connection).ReadLine()
		if err != nil {
			log.Fatal("Error reading from buffer\n", err)
		}
		data, err := parseMessage(string(message))
		if err != nil {
			log.Fatal("Error parsing message\n", err)
		}

		fmt.Println("Tags:", data.Tags)
		fmt.Println("Source:", data.Source)
		fmt.Println("Command:", data.Command)
		fmt.Println("Parameters:", data.Parameters)

		fmt.Println((string(message)))
	}
}
