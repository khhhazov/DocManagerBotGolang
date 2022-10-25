package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateRoleLevelKeyboard() tgbotapi.InlineKeyboardMarkup {
	var roleLevelKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Министерство образования",
				"Министерство"),
			tgbotapi.NewInlineKeyboardButtonData("Ректор университета",
				"Ректор"),
			tgbotapi.NewInlineKeyboardButtonData("Президент университета",
				"Президент"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проректор",
				"Проректор"),
			tgbotapi.NewInlineKeyboardButtonData("Директор",
				"Директор"),
			tgbotapi.NewInlineKeyboardButtonData("Заведущий кафедры",
				"Заведущий"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Преподаватель",
				"Преподаватель"),
			tgbotapi.NewInlineKeyboardButtonData("Староста учебной группы",
				"Староста"),
			tgbotapi.NewInlineKeyboardButtonData("Студент",
				"Студент"),
		),
	)

	return roleLevelKeyboard
}
