package sqllite

import (
	"GameServer/utils/helper"
	"database/sql"
	"log"
	"os"
)

func GetSqlInstance() (*sql.DB, error) {
	if helper.DoesFileExist("levelsdb.db") {
		return sql.Open("sqlite3", "./levelsdb.db")
	}

	file, err := os.Create("levelsdb.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	return sql.Open("sqlite3", "./levelsdb.db")
}
