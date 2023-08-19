package main

import (
	"fmt"
	"net"
	"strings"

	"ryanwarsaw.com/espresso/protocol"
	"ryanwarsaw.com/espresso/ui"
)

func PingEventHandler(message protocol.Message, conn net.Conn) {
	// handle optional token parameter found on more modern server architectures.
	if len(message.Parameters) > 0 {
		conn.Write([]byte(fmt.Sprintf("PONG %s \r\n", message.Parameters[0])))
	} else {
		conn.Write([]byte("PONG\r\n"))
	}
}

func CapEventHandler(message protocol.Message, conn net.Conn) {
	conn.Write([]byte("CAP END\r\n"))
}

func JoinChannelEventHandler(message protocol.Message, conn net.Conn) {
	if message.Source.Nick == *options.Username {
		channelName := message.Parameters[0][1:]
		ui.GetApplication().QueueUpdateDraw(func() {
			ui.GetChannelPanel().AddItem(channelName, "", 0, nil)
		})
	} else {
		ui.GetApplication().QueueUpdateDraw(func() {
			formatted := fmt.Sprintf("[darkseagreen]%s has joined the channel", message.Source.Nick)
			ui.GetMessagePanel().AddItem(formatted, "", 0, nil)
		})
	}
}

func PartChannelEventHandler(message protocol.Message, conn net.Conn) {
	if message.Source.Nick == *options.Username {
		channelName := message.Parameters[0][1:]
		indices := ui.GetChannelPanel().FindItems(channelName, "", true, true)
		if len(indices) > 0 {
			ui.GetApplication().QueueUpdateDraw(func() {
				ui.GetChannelPanel().RemoveItem(indices[0])
			})
		}
	} else {
		ui.GetApplication().QueueUpdateDraw(func() {
			formatted := fmt.Sprintf("[orange]%s has left the channel", message.Source.Nick)
			ui.GetMessagePanel().AddItem(formatted, "", 0, nil)
		})
	}
}

func MessageEventHandler(message protocol.Message, conn net.Conn) {
	sanitized := strings.Join(message.Parameters[1:], " ")
	if sanitized[0] == ':' {
		sanitized = sanitized[1:]
	}
	ui.GetApplication().QueueUpdateDraw(func() {
		formatted := fmt.Sprintf("[darkseagreen]%s: [white]%s", message.Source.Nick, sanitized)
		ui.GetMessagePanel().AddItem(formatted, "", 0, nil)
	})
}

func RPLEventHandler(message protocol.Message, conn net.Conn) {
	sanitized := strings.Join(message.Parameters[1:], " ")
	if sanitized[0] == ':' {
		sanitized = sanitized[1:]
	}
	ui.GetApplication().QueueUpdateDraw(func() {
		formatted := fmt.Sprintf("[darkseagreen]%s", sanitized)
		ui.GetMessagePanel().AddItem(formatted, "", 0, nil)
	})
}

func EventDispatcher(message protocol.Message, conn net.Conn) {
	switch command := message.Command; command {
	case "PING":
		PingEventHandler(message, conn)
	case "CAP":
		CapEventHandler(message, conn)
	case "JOIN":
		JoinChannelEventHandler(message, conn)
	case "PART":
		PartChannelEventHandler(message, conn)
	case "PRIVMSG":
		MessageEventHandler(message, conn)
	case "001", "002", "003":
		RPLEventHandler(message, conn)
	}
}
