package event

import (
	"sync"

	"github.com/4mti/ponto/domain/core"
)

type EventListenerImpl struct {
	mu        sync.RWMutex
	listeners map[string][]chan core.Event
}

func (this *EventListenerImpl) Emit(eventName string, data interface{}) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	channels, exists := this.listeners[eventName]

	if !exists || len(channels) == 0 {
		return nil
	}

	event := core.Event{
		Name: eventName,
		Data: data,
	}

	for _, channel := range channels {
		select {
		case channel <- event:
		default:
		}
	}

	return nil
}

func (this *EventListenerImpl) On(eventName string) <-chan core.Event {
	this.mu.Lock()
	defer this.mu.Unlock()

	eventChan := make(chan core.Event)
	this.listeners[eventName] = append(this.listeners[eventName], eventChan)

	return eventChan
}

func (this *EventListenerImpl) Clear() {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, chans := range this.listeners {
		for _, channel := range chans {
			close(channel)
		}
	}

	this.listeners = make(map[string][]chan core.Event)
}

func NewEventListener() core.EventListener {
	return &EventListenerImpl{
		listeners: make(map[string][]chan core.Event),
	}
}
