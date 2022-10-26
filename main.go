package main

import (
	"DocManagerBot/models"
)

import (
	"DocManagerBot/factories"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func main() {
	var updates = createConfig()
	var ending = false
	var info = models.Info{Role: "",
		TaskType: "", Weight: 0}

	for update := range updates {
		if update.Message != nil {
			var command = factories.FactoryCommands(update, &ending)
			var msg = factories.FactoryMsg(command, update, &info)
			if msg.Text != "" {
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
			}
			log.Printf("[%s, %s, %d]", info.Role, info.TaskType, info.Weight)
		} else if update.CallbackQuery != nil {
			var msg = factories.FactoryAction(update, &ending, &info)
			if msg.Text != "" {
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		}
	}
}
