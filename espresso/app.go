package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"ryanwarsaw.com/protocol"
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
		data, err := protocol.ParseMessage(string(message))
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
