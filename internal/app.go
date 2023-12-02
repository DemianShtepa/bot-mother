package internal

type Applications map[string]Application

type Application interface {
	Process(chan Event)
	GetName() string
}

func (a Applications) Process() Events {
	events := make(chan Event)

	for _, application := range a {
		go application.Process(events)
	}

	return events
}
