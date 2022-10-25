package factories

import (
	"DocManagerBot/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func FactoryMsg(command string, update tgbotapi.Update) tgbotapi.MessageConfig {
	var msg = tgbotapi.NewMessage(update.Message.Chat.ID,
		"")

	switch command {
	case "/start":
		msg.Text = "Здравствуйте и добро пожаловать в \"Документовед\"! " +
			"Для продолжения выберите один из пунктов:"
		msg.ReplyMarkup = keyboards.CreateRoleLevelKeyboard()
		return msg
	case "/date":
		msg.Text = "Поздравляю!"
		return msg
	default:
		return msg
	}
}
