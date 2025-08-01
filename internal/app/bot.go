package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Bot(d Dependencies) {
	states := make(map[int64]string)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := d.BotAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			text := update.Message.Text

			switch {
			case text == "/start":
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Погода"),
					),
				)

				msg := tgbotapi.NewMessage(chatID, "Привет!")
				msg.ReplyMarkup = keyboard

				if _, err := d.BotAPI.Send(msg); err != nil {
					fmt.Println("Ошибка при отправке кнопок:", err)
				}
			case states[chatID] == "city_await":
				fmt.Println(text)
				res, err := d.OW.Get(text)
				if err != nil {
					fmt.Println("Error to get weather:", err)
				}
				d.BotAPI.Send(tgbotapi.NewMessage(chatID, res))
				delete(states, chatID)
			case text == "Погода":
				_, err := d.BotAPI.Send(tgbotapi.NewMessage(chatID, "Введите город:"))
				if err != nil {
					fmt.Println("Send message error:", err)
				}
				states[chatID] = "city_await"
			default:
				if _, err := d.BotAPI.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда")); err != nil {
					fmt.Println("Send message error:", err)
				}
			}
		}
	}
}
