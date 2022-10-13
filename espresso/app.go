package main

import (
	"bufio"
	"fmt"
	"github.com/rivo/tview"
	"log"
	"net"
	"os"
	"ryanwarsaw.com/protocol"
)

var messagePanel = MessagePanel()
var inputPanel = InputPanel()
var messageLayout = MessageLayout(messagePanel, inputPanel)

var channelPanel = ChannelPanel()
var userPanel = UserPanel()
var sidebarLayout = SideBarLayout(channelPanel, userPanel)

var appLayout = AppLayout(messageLayout, sidebarLayout)

func CreateDebugFile() *os.File {
	file, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error creating debug file: %v", err)
	}
	return file
}

func RenderApplication() {
	app := tview.NewApplication()
	app.SetRoot(appLayout, true)
	app.SetFocus(inputPanel)
	if err := app.Run(); err != nil {
		log.Fatalf("Error rendering application: %v", err)
	}
	defer os.Exit(0)
}

func main() {
	options := ConfigureAndParseFlags()

	go RenderApplication()

	if *options.Debug {
		file := CreateDebugFile()
		defer file.Close()
		log.SetOutput(file)
	}

	log.Printf("Connecting to server with options:\n%+v\n", &options)
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
		messagePanel.AddItem(string(message), "", 0, nil)

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
