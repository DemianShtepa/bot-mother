package ukrzaliznytsia

import (
	"bot-mother/internal"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Application struct {
	env Env
}

func NewApplication(env Env) *Application {
	return &Application{env: env}
}

func (a Application) Process(events chan internal.Event) {
	request, err := http.NewRequest("GET", a.env.ApiURL, nil)
	if err != nil {
		fmt.Println("Error while creating request: ", err)
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "UZ/1.10.1 iOS/17.1.1 User/1374232")
	request.Header.Add("Authorization", "Bearer "+a.env.ApiToken)

	client := http.Client{}

	query := request.URL.Query()
	query.Add("station_from_id", strconv.Itoa(a.env.StationFrom))
	query.Add("station_to_id", strconv.Itoa(a.env.StationTo))
	query.Add("with_transfers", "0")

	for {
		for date := a.env.DateFrom; date.Before(a.env.DateTo); date = date.AddDate(0, 0, 1) {
			query.Add("date", date.Format("2006-01-02"))
			request.URL.RawQuery = query.Encode()
			response, err := client.Do(request)
			if err != nil {
				fmt.Println("Error while performing request: ", err)
				break
			}

			defer response.Body.Close()

			tripsResponse := trips{}
			if jsonDecoder := json.NewDecoder(response.Body); jsonDecoder.Decode(&tripsResponse) == nil {
				if len(tripsResponse.Direct) > 1 {
					events <- Event{date: date}
				}
			}

			time.Sleep(time.Second * 10)
		}

		time.Sleep(time.Minute * 30)
	}
}
