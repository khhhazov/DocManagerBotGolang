package factories

import (
	"DocManagerBot/calculation"
	db2 "DocManagerBot/db"
	"DocManagerBot/keyboards"
	"DocManagerBot/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func FactoryMsg(command string, update tgbotapi.Update,
	info *models.Info,
) tgbotapi.MessageConfig {
	var msg = tgbotapi.NewMessage(update.Message.Chat.ID,
		"")

	switch command {
	case "/start":
		msg.Text = "Здравствуйте и добро пожаловать в \"Документовед\"! " +
			"Для продолжения выберите один из пунктов:"
		msg.ReplyMarkup = keyboards.CreateRoleLevelKeyboard()
		return msg
	case "/date":
		var dt, err = time.Parse("02.01.2006", update.Message.Text)
		if err != nil {
			panic(err)
		}

		/*DATE*/
		crntDate := time.Now()
		cmplDate := dt
		rsltDateDays := (cmplDate.Sub(crntDate) / time.Minute) / 1200

		/*CULC_FUNCTION*/
		role := calculation.CalcRoleWeight(info.Role)
		taskType := calculation.CalcTaskTypeWeight(info.TaskType)
		ki := role + taskType + int(rsltDateDays)
		info.Weight = ki * ((1 * int(rsltDateDays)) + (2 * role) + (3 * taskType))

		db := db2.ConnectDB()
		db2.InsertWeight(info.Weight, info.Role,
			info.TaskType, crntDate.Format("02.01.2006"),
			cmplDate.Format("02.01.2006"), db)

		msg.Text = fmt.Sprintf("Спасибо за уделенное время\nЗадача находится на рассмотрении\n Weight:[%d]", info.Weight)
		return msg
	default:
		return msg
	}
}
