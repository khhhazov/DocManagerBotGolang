package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateTaskTypeKeyboard() tgbotapi.InlineKeyboardMarkup {
	var taskTypeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Учебная",
				"Учебная"),
			tgbotapi.NewInlineKeyboardButtonData("Научная",
				"Научная"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Административная",
				"Административная"),
			tgbotapi.NewInlineKeyboardButtonData("Культурно массовая",
				"Культурно массовая"),
		),
	)

	return taskTypeKeyboard
}
