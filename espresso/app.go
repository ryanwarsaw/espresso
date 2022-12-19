package main

import (
	"bufio"
	"fmt"
	"github.com/rivo/tview"
	"log"
	"net"
	"os"
	"ryanwarsaw.com/protocol"
	"github.com/gdamore/tcell/v2"
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

func main() {
	options := ConfigureAndParseFlags()

	if *options.Debug {
		file := CreateDebugFile()
		defer file.Close()
		log.SetOutput(file)
	}

	app := tview.NewApplication()
	app.SetRoot(appLayout, true)
	app.SetFocus(inputPanel)

	// Akin to a UI thread, prevents the renderer and event loop
	// from blocking the TCP client listener
	go func() {
		if err := app.Run(); err != nil {
			log.Fatalf("Error rendering application: %v", err)
		}
		defer os.Exit(0)
	}()

	log.Printf("Connecting to server with options:\n%+v\n", &options)
	hostAddress := fmt.Sprintf("%s:%d", *options.Address, *options.Port)

	connection, err := net.Dial("tcp4", hostAddress)
	if err != nil {
		log.Fatal("Failed to establish connection to host server\n", err)
	}

	connection.Write(protocol.Commands.CapList())
	connection.Write(protocol.Commands.Nick(*options.Username))
	connection.Write(protocol.Commands.User(*options.Username))

	inputPanel.SetDoneFunc(func(key tcell.Key) {
		if (key == tcell.KeyEnter) {
			message := inputPanel.GetText()
			connection.Write([]byte("PRIVMSG #foobaz :" + message + "\r\n"))
			messagePanel.AddItem(*options.Username + ": " + (message), "", 0, nil)
			inputPanel.SetText("")
		}
	})

	buffer := bufio.NewReader(connection)
	for {
		message, _, err := buffer.ReadLine()
		if err != nil {
			log.Fatal("Error reading from buffer\n", err)
		}
		log.Println((string(message)))
		app.QueueUpdateDraw(func() {
			messagePanel.AddItem(string(message), "", 0, nil)
		})

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
