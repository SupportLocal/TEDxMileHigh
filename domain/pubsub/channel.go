package pubsub

type Channel string

const (
	MessageBlocked Channel = "message-blocked"
	MessageCreated Channel = "message-created"
	MessageCycled  Channel = "message-cycled"
	MessageSaved   Channel = "message-saved"
	MessageUpdated Channel = "message-updated"
)

func ChannelFor(channelName string) Channel {
	switch channelName {
	case "message-blocked":
		return MessageBlocked
	case "message-created":
		return MessageCreated
	case "message-cycled":
		return MessageCycled
	case "message-saved":
		return MessageSaved
	case "message-updated":
		return MessageUpdated
	default:
		panic("invalid channel")
	}
}

func (c Channel) String() string {
	return string(c)
}
