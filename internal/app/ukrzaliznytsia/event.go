package ukrzaliznytsia

import "time"

type Event struct {
	date time.Time
}

func (e Event) GetMessage() string {
	return "З'явилось мiсце - " + e.date.Format("2006-01-02")
}
