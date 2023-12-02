package internal

type Events chan Event

func (e Events) Notify(bot Bot) {
	for event := range e {
		bot.Notify(event)
	}
}

type Event interface {
	GetMessage() string
}
