package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Env struct {
	BotToken string
	ChatId   int64
}

func NewEnv(
	botToken string,
	chatId int64,
) *Env {
	return &Env{
		BotToken: botToken,
		ChatId:   chatId,
	}
}

type Bot struct {
	telegramBot *tgbotapi.BotAPI
	chatId      int64
}

func NewBot(env Env) *Bot {
	telegramBot, err := tgbotapi.NewBotAPI(env.BotToken)
	if err != nil {
		log.Fatal("Error for creating bot")
	}

	return &Bot{telegramBot: telegramBot, chatId: env.ChatId}
}

func (bot Bot) Notify(event Event) {
	message := tgbotapi.NewMessage(bot.chatId, event.GetMessage())

	if _, err := bot.telegramBot.Send(message); err != nil {
		log.Fatal(err)
	}
}
