package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"ryanwarsaw.com/protocol"
)

type ConnectionOptions struct {
	Address  *string
	Port     *int
	Username *string
	Debug    *bool
}

func ConfigureAndParseFlags() ConnectionOptions {
	addressPtr := flag.String("address", "localhost", "IRC server address")
	portPtr := flag.Int("port", 6667, "IRC server port")
	usernamePtr := flag.String("username", "ryan", "Personal identifier")
	debugPtr := flag.Bool("debug", false, "Enables network logging to debug.log file")
	flag.Parse()
	return ConnectionOptions{
		Address:  addressPtr,
		Port:     portPtr,
		Username: usernamePtr,
		Debug:    debugPtr,
	}
}

func CreateDebugFile() *os.File {
	file, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error creating debug file: %v", err)
	}
	return file
}

func main() {
	options := ConfigureAndParseFlags()

	if *options.Debug {
		file := CreateDebugFile()
		defer file.Close()
		log.SetOutput(file)
	}

	fmt.Println("Connecting to server with address:", *options.Address)
	fmt.Println("Using port:", *options.Port)
	fmt.Println("Using username:", *options.Username)

	hostAddress := fmt.Sprintf("%s:%d", *options.Address, *options.Port)

	connection, err := net.Dial("tcp4", hostAddress)
	if err != nil {
		log.Fatal("Failed to establish connection to host server\n", err)
	}

	connection.Write(protocol.Commands.CapList())
	connection.Write(protocol.Commands.Nick(*options.Username))
	connection.Write(protocol.Commands.User(*options.Username))

	buffer := bufio.NewReader(connection)
	for {
		message, _, err := buffer.ReadLine()
		if err != nil {
			log.Fatal("Error reading from buffer\n", err)
		}
		log.Println((string(message)))

		data, err := protocol.ParseMessage(string(message))
		if err != nil {
			log.Fatal("Error parsing message\n", err)
		}

		// TODO: Move this to a dispatcher
		if data.Command == "PING" {
			connection.Write(protocol.Responses.PingResponseMessage(data))
		}

		// TODO: Handle state of cap life cycle instead of this
		if data.Command == "CAP" {
			connection.Write(protocol.Commands.CapEnd())
			connection.Write([]byte("PING test\r\n"))
			connection.Write([]byte("JOIN #foobaz\r\n"))
		}
	}
}
