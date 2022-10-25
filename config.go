package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var bot, err = tgbotapi.NewBotAPI("5677711562:AAHJ5axUEveXPtnJrxTYT7U4oEEY35ZFGSo")

func createConfig() tgbotapi.UpdatesChannel {
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var updateConfig = tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return bot.GetUpdatesChan(updateConfig)
}
