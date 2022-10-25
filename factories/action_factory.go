package factories

import (
	"DocManagerBot/keyboards"
	"DocManagerBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func FactoryAction(update tgbotapi.Update, ending *bool, info *models.Info) tgbotapi.MessageConfig {

	var msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
		"")
	switch update.CallbackQuery.Message.Text {
	case "Здравствуйте и добро пожаловать в \"Документовед\"! " +
		"Для продолжения выберите один из пунктов:":
		info.Role = update.CallbackQuery.Data
		msg.Text = "Отлично! Теперь выберите тип задачи:"
		msg.ReplyMarkup = keyboards.CreateTaskTypeKeyboard()
		return msg
	case "Отлично! Теперь выберите тип задачи:":
		info.TaskType = update.CallbackQuery.Data
		msg.Text = "Введите дату завершения задачи в формате dd.mm.yyyy:"
		*ending = true
		return msg
	default:
		return msg
	}
}
