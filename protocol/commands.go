package protocol

import "fmt"

var Commands commands

type commands struct{}

func (commands) CapList() []byte {
	return []byte("CAP LS 302\r\n")
}

func (commands) CapEnd() []byte {
	return []byte("CAP END\r\n")
}

func (commands) Nick(nickname string) []byte {
	return []byte(fmt.Sprintf("NICK %s \r\n", nickname))
}

func (commands) User(username string) []byte {
	return []byte(fmt.Sprintf("USER %s 0 * :%s\r\n", username, username))
}
