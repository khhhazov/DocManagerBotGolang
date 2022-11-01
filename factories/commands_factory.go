package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
)

func FactoryCommands(update tgbotapi.Update, ending *bool) string {
	if *ending {
		matched, _ := regexp.MatchString(`\d\d[.]\d\d[.]\d\d\d\d`, update.Message.Text)
		if matched {
			*ending = false
			return "/date"
		}
	}
	switch update.Message.Text {
	case "/start":
		return update.Message.Text
	case "/clear":
		return update.Message.Text
	default:
		return ""
	}
}
