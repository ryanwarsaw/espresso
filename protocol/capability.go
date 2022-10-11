package protocol

type Capability string

// All IRC v3 spec server negotiable capabilities
// See: https://ircv3.net/registry#capabilities
var Capabilities = struct {
	AccountNotify   Capability
	AccountTag      Capability
	AwayNotify      Capability
	Batch           Capability
	ChgHost         Capability
	EchoMessage     Capability
	ExtendedJoin    Capability
	InviteNotify    Capability
	MessageTags     Capability
	Monitor         Capability
	MultiPrefix     Capability
	Sasl            Capability
	ServerTime      Capability
	SetName         Capability
	UserHostInNames Capability
}{
	AccountNotify:   "account-notify",
	AccountTag:      "account-tag",
	AwayNotify:      "away-notify",
	Batch:           "batch",
	ChgHost:         "chghost",
	EchoMessage:     "echo-message",
	ExtendedJoin:    "extended-join",
	InviteNotify:    "invite-notify",
	MessageTags:     "message-tags",
	Monitor:         "monitor",
	MultiPrefix:     "multi-prefix",
	Sasl:            "sasl",
	ServerTime:      "server-time",
	SetName:         "setname",
	UserHostInNames: "userhost-in-names",
}
