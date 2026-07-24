package core

type Event struct {
	Name string
	Data interface{}
}

type EventListener interface {
	Emit(eventName string, data interface{}) error
	On(eventName string) <-chan Event
	Clear()
}
