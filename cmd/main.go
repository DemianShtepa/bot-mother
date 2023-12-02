package main

import (
	"bot-mother/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
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

func registerApplications() internal.Applications {
	applications := make(map[string]internal.Application)

	return applications
}

func main() {
	env := parseEnv()

	applications := registerApplications()
	bot := internal.NewBot(*env)

	events := applications.Process()
	events.Notify(*bot)
}
