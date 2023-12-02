package app

import (
	"bot-mother/internal"
)

type Applications []Application

type Application interface {
	Process(chan internal.Event)
}

func (a Applications) Process() internal.Events {
	events := make(chan internal.Event)

	for _, application := range a {
		go application.Process(events)
	}

	return events
}
