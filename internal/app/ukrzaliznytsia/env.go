package ukrzaliznytsia

import "time"

type Env struct {
	ApiURL   string
	ApiToken string

	StationFrom int
	StationTo   int
	DateFrom    time.Time
	DateTo      time.Time
}

func NewEnv(
	apiURL string,
	apiToken string,
	stationFrom int,
	stationTo int,
	dateFrom time.Time,
	dateTo time.Time,
) *Env {
	return &Env{
		ApiURL:      apiURL,
		ApiToken:    apiToken,
		StationFrom: stationFrom,
		StationTo:   stationTo,
		DateFrom:    dateFrom,
		DateTo:      dateTo,
	}
}
