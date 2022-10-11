package protocol

import (
	"reflect"
	"testing"
)

func TestParseMessageEmpty(t *testing.T) {
	_, err := ParseMessage("")
	if err == nil || err.Error() != "empty message" {
		t.Fatalf("Should throw error when parsing empty message")
	}
}

func TestParseMessagePlain(t *testing.T) {
	expect := Message{
		Tags:       make(map[string]string),
		Source:     "",
		Command:    "CAP",
		Parameters: []string{"LS", "*", ":multi-prefix", "extended-join", "sasl"},
	}
	message, err := ParseMessage("CAP LS * :multi-prefix extended-join sasl")
	if !reflect.DeepEqual(message, expect) || err != nil {
		t.Fatalf("Should handle plain message")
	}
}

func TestParseMessageSource(t *testing.T) {
	expect := Message{
		Tags:       make(map[string]string),
		Source:     "irc.example.com",
		Command:    "CAP",
		Parameters: []string{"LS", "*", ":multi-prefix", "extended-join", "sasl"},
	}
	message, err := ParseMessage(":irc.example.com CAP LS * :multi-prefix extended-join sasl")
	if !reflect.DeepEqual(message, expect) || err != nil {
		t.Fatalf("Should detect optional message source field")
	}
}

func TestParseMessageTags(t *testing.T) {
	expect := Message{
		Tags: map[string]string{
			"id":        "234AB",
			"custom/id": "metadata",
		},
		Source:     "",
		Command:    "PRIVMSG",
		Parameters: []string{"#chan", ":Hey", "what's", "up!"},
	}
	message, err := ParseMessage("@id=234AB;custom/id=metadata PRIVMSG #chan :Hey what's up!")
	if !reflect.DeepEqual(message, expect) || err != nil {
		t.Fatalf("Should detect optional message tag field")
	}
}

func TestParseMessageInvalidTag(t *testing.T) {
	_, err := ParseMessage("@id=234AB= PRIVMSG #chan :Hey what's up!")
	if err == nil || err.Error() != "invalid tag format" {
		t.Fatalf("Should throw error when parsing invalid tag")
	}
}

func TestParseMessageSourceAndTags(t *testing.T) {
	expect := Message{
		Tags: map[string]string{
			"id":        "234AB",
			"custom/id": "metadata",
		},
		Source:     "irc.example.com",
		Command:    "PRIVMSG",
		Parameters: []string{"#chan", ":Hey", "what's", "up!"},
	}
	message, err := ParseMessage("@id=234AB;custom/id=metadata :irc.example.com PRIVMSG #chan :Hey what's up!")
	if !reflect.DeepEqual(message, expect) || err != nil {
		t.Fatalf("Should detect both source and tags")
	}
}
