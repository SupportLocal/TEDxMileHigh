package current_message

import (
	eventsource "github.com/antage/eventsource/http"
)

func NewEventSource() eventsource.EventSource {
	es := eventsource.New(nil)

	return es
}
