package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_weather/pkg/ow"
)

type UseCase struct {
	OW     *ow.OW
	BotAPI *tgbotapi.BotAPI
	States map[int64]string
}

func (u *UseCase) SendStart(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Погода"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Привет!")
	msg.ReplyMarkup = keyboard

	if _, err := u.BotAPI.Send(msg); err != nil {
		fmt.Println("Ошибка при отправке кнопок:", err)
	}
}

func (u *UseCase) SendWeather(update tgbotapi.Update) {
	text := update.Message.Text
	chatID := update.Message.Chat.ID

	fmt.Println(text)
	res, err := u.OW.Get(text)
	if err != nil {
		fmt.Println("Error to get weather:", err)
	}
	_, err = u.BotAPI.Send(tgbotapi.NewMessage(chatID, res))
	if err != nil {
		fmt.Println("Error to send weather:", err)
	}
	delete(u.States, chatID)
}

func (u *UseCase) CityAwait(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	_, err := u.BotAPI.Send(tgbotapi.NewMessage(chatID, "Введите город:"))
	if err != nil {
		fmt.Println("Send message error:", err)
	}
	u.States[chatID] = "city_await"
}

func Bot(d Dependencies) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	uc := UseCase{OW: d.OW, BotAPI: d.BotAPI.BotAPI, States: make(map[int64]string)}

	updates := d.BotAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			text := update.Message.Text

			switch {
			case text == "/start":
				uc.SendStart(update)
			case uc.States[chatID] == "city_await":
				uc.SendWeather(update)
			case text == "Погода":
				uc.CityAwait(update)
			default:
				if _, err := d.BotAPI.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда")); err != nil {
					fmt.Println("Send message error:", err)
				}
			}
		}
	}
}
