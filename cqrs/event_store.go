package cqrs

//Represents a backing store to save aggregate's events
type EventStore interface {
	GetAllEvents() []Event
	GetEvents(aggregateId string) []Event
	GetEventsFrom(aggregateName string, offset, count int) []Event
	GetAllEventsFrom(offset, count int) []Event
	GetEvent(id string) Event
	SaveEvents(aggregateId string, events []Event) error
}