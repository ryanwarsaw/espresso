package main

import (
	"flag"
	"fmt"
)

type ConnectionOptions struct {
	Address  *string
	Port     *int
	Username *string
	Debug    *bool
}

// Handles serialization of the ConnectionOptions object to SprintF
func (options *ConnectionOptions) String() string {
	return fmt.Sprintf("ConnectionOptions{Address:%s, Port:%d, Username:%s, Debug:%t}",
		*options.Address, *options.Port, *options.Username, *options.Debug)
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
