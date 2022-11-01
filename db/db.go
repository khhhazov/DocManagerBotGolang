package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func ConnectDB() *sqlx.DB {
	connStr := "host=178.20.47.138 port=5432 user=vymanvar password=1234 dbname=docdb sslmode=disable"
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error connecting db: %s", err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func InsertWeight(weight int, role string,
	task_type string, crnt_time string,
	compl_time string, db *sqlx.DB,
) {
	query := fmt.Sprintf("INSERT INTO weight_table (weight, role, task_type, crnt_time, compl_time) values (%d, '%s', '%s', '%s', '%s')", weight, role, task_type, crnt_time, compl_time)
	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}

func ClearDataBase(db *sqlx.DB) {
	query := "DELETE FROM weight_table"
	_, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	_, err = db.Query("ALTER SEQUENCE weight_table_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}
