package protocol

import "fmt"

func CapListCommand() []byte {
	return []byte("CAP LS 302\r\n")
}

func NickCommand(nickname string) []byte {
	return []byte(fmt.Sprintf("NICK %s \r\n", nickname))
}

func UserCommand(username string) []byte {
	return []byte(fmt.Sprintf("USER %s 0 * :%s\r\n", username, username))
}