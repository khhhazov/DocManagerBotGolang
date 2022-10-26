package main

import (
	"DocManagerBot/factories"
	"DocManagerBot/models"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func main() {
	var updates = createConfig()
	var ending = false
	var info = models.Info{Role: "",
		TaskType: "", Weight: 0}

	connStr := "host=178.20.47.138 port=5432 user=vymanvar password=1234 dbname=docdb sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into weight_table (weight, role, task_type, crnt_time, compl_time) values (194, 'Вася', 'Курсовая', '26.10.2022', '30.10.2022')")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected())

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
