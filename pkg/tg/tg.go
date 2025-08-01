package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	Token string `envconfig:"BOT_TOKEN" required:"true"`
}
type Bot struct {
	*tgbotapi.BotAPI
}

func NewClient(c Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	return &Bot{bot}, nil
}
