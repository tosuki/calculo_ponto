package event

import "github.com/4mti/ponto/domain/core"

type EventListenerImpl struct {
}

func Emit(eventName string, data interface{}) error {
	return nil
}

func On(eventName string) <-chan core.Event {
	eventChan := make(chan core.Event)

	return eventChan
}

func Clear() {
	//...
}
