package main

import (
	"DocManagerBot/factories"
)

func main() {
	var updates = createConfig()
	var ending = false
	//var info = models.Info{Role: 0, TaskType: 0, Weight: 0}

	for update := range updates {
		if update.Message != nil {
			var command = factories.FactoryCommands(update, &ending)
			var msg = factories.FactoryMsg(command, update)
			if msg.Text != "" {
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
			}
		} else if update.CallbackQuery != nil {
			/*после тогда как нажата кнопка (наважно какая),
			происходит продолжение команд*/

			var msg = factories.FactoryAction(update, &ending)
			if msg.Text != "" {
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		}
	}
}
