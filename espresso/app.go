package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"ryanwarsaw.com/espresso/protocol"
	"ryanwarsaw.com/espresso/ui"
)

func JoinChannel(channelName string, connection net.Conn) {
	connection.Write([]byte(fmt.Sprintf("JOIN %s\r\n", channelName)))
}

func LeaveChannel(channelName string, connection net.Conn) {
	connection.Write([]byte(fmt.Sprintf("PART %s\r\n", channelName)))
}

func render() {
	if err := ui.GetApplication().Run(); err != nil {
		log.Fatalf("Error rendering application: %v", err)
	}
	defer os.Exit(0)
}

var options ConnectionOptions

func main() {
	options = ConfigureAndParseFlags()

	if *options.Debug {
		file := CreateDebugFile()
		defer file.Close()
		log.SetOutput(file)
	}
	ui.GetApplication().SetRoot(ui.GetAppLayout(), true)
	ui.GetApplication().SetFocus(ui.GetInputPanel())

	// Akin to a UI thread, prevents the renderer and event loop
	// from blocking the TCP client listener
	go render()

	log.Printf("Connecting to server with options:\n%+v\n", &options)
	hostAddress := fmt.Sprintf("%s:%d", *options.Address, *options.Port)

	connection, err := net.Dial("tcp4", hostAddress)
	if err != nil {
		log.Fatal("Failed to establish connection to host server\n", err)
	}

	connection.Write(protocol.CapListCommand())
	connection.Write(protocol.NickCommand(*options.Username))
	connection.Write(protocol.UserCommand(*options.Username))

	ui.GetInputPanel().SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			message := ui.GetInputPanel().GetText()

			if len(message) > 0 {
				if message[0] == '/' {
					segments := strings.Split(message[1:], " ")
					if len(segments) == 2 && strings.ToLower(segments[0]) == "join" {
						channelName := strings.TrimSpace(strings.ToLower(segments[1]))
						JoinChannel(channelName, connection)
					}

					if len(segments) == 2 && strings.ToLower(segments[0]) == "leave" {
						channelName := strings.TrimSpace(strings.ToLower(segments[1]))
						LeaveChannel(channelName, connection)
					}
				} else {
					connection.Write([]byte("PRIVMSG #foobaz :" + message + "\r\n"))
					ui.GetMessagePanel().AddItem("[darkseagreen]"+*options.Username+": [white]"+(message), "", 0, nil)
				}

				ui.GetInputPanel().SetText("")
			}
		}
	})

	buffer := bufio.NewReader(connection)
	for {
		message, _, err := buffer.ReadLine()
		if err != nil {
			log.Fatal("Error reading from buffer\n", err)
		}

		payload, err := protocol.ParseMessage(string(message))
		if err != nil {
			log.Fatal("Error parsing message\n", err)
		}

		log.Println(payload)
		EventDispatcher(payload, connection)
	}
}
