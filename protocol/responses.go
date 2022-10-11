package protocol

import "fmt"

var Responses responses

type responses struct{}

// Generates the response to an IRC server Ping handling the
// optional token parameter found on more modern server architectures.
func (responses) PingResponseMessage(message Message) []byte {
	if len(message.Parameters) > 0 {
		return []byte(fmt.Sprintf("PONG %s \r\n", message.Parameters[0]))
	}
	return []byte("PONG\r\n")
}
