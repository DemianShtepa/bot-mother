package main

import (
	"bot-mother/internal"
	"bot-mother/internal/app"
	"bot-mother/internal/app/ukrzaliznytsia"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func parseEnv() *internal.Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatId, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	return internal.NewEnv(
		os.Getenv("BOT_TOKEN"),
		chatId,
	)
}

func registerApplications() app.Applications {
	var applications app.Applications

	applications = append(applications, ukrzaliznytsia.NewApplication(*func() *ukrzaliznytsia.Env {
		stationFrom, _ := strconv.Atoi(os.Getenv("STATION_FROM"))
		stationTo, _ := strconv.Atoi(os.Getenv("STATION_TO"))

		dateFrom, _ := time.Parse("2006-01-02", os.Getenv("DATE_FROM"))
		dateTo, _ := time.Parse("2006-01-02", os.Getenv("DATE_TO"))

		return ukrzaliznytsia.NewEnv(
			os.Getenv("API_URL"),
			os.Getenv("API_TOKEN"),
			stationFrom,
			stationTo,
			dateFrom,
			dateTo,
		)
	}()))

	return applications
}

func main() {
	env := parseEnv()

	applications := registerApplications()
	bot := internal.NewBot(*env)

	events := applications.Process()
	events.Notify(*bot)
}
