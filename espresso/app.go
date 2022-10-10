package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"bufio"
	"ryanwarsaw.com/protocol"
)

type ConnectionOptions struct {
	Address *string
	Port *int
	Username *string
}

func ConfigureAndParseFlags() (ConnectionOptions) {
	addressPtr := flag.String("address", "localhost", "IRC server address")
	portPtr := flag.Int("port", 6667, "IRC server port")
	usernamePtr := flag.String("username", "ryan", "Personal identifier")
	flag.Parse()
	return ConnectionOptions {
		Address: addressPtr,
		Port: portPtr,
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

	status, _, err := bufio.NewReader(connection).ReadLine()
	if err != nil {
		log.Fatal("Error creating buffered reader")
	}
	fmt.Println(string(status))
	fmt.Println(protocol.Capability("account-notify") == protocol.Capabilities.AccountNotify)
}