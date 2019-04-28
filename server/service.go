package main

import "bytes"

type PingCreatedEvent struct {
	Encoder
	Hop int
}

type Service struct {
	logger *Logger
	bus    *Bus
	name   string
}

func NewService(logger *Logger, bus *Bus, serviceName string) *Service {
	return &Service{
		logger: logger,
		bus:    bus,
		name:   serviceName,
	}
}

func (service *Service) CreatePing(lastHop int) error {
	event := newPingCreatedEvent(lastHop)
	data, err := event.Encode()
	if err != nil {
		return err
	}
	return service.bus.Send("PingCreatedEvent", data)
}

func newPingCreatedEvent(lastHop int) *PingCreatedEvent {
	event := &PingCreatedEvent{
		Hop: lastHop + 1,
	}
	return event
}

func (event *PingCreatedEvent) Encode() ([]byte, error) {
	b := &bytes.Buffer{}
	err := event.encode(b, event)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
