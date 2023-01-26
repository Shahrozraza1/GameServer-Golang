package db

import (
	"GameServer/consts"
	"GameServer/db/sqllite"
	"GameServer/model"
	"GameServer/utils/helper"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Manager struct {
	db *sql.DB
}

var manger *Manager

func GetManagerInstance() *Manager {
	if manger == nil {
		manger = &Manager{}
		manger.Init()
		return manger
	}
	return manger
}

func (manager *Manager) Init() {
	var err error
	manager.db, err = sqllite.GetSqlInstance()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (manager *Manager) CreateTables() {
	manager.Init()
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS levels (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"level" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Creating Levels table...")
	statement, err := manager.db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Levels table created")
	manager.db.Close()
}

func (manager *Manager) AddLevelToDatabase(level *model.Level) {
	manager.Init()
	defer manager.db.Close()

	insertStudentSQL := `INSERT INTO levels(id, level) VALUES (?, ?)`
	statement, err := manager.db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(level.ID, helper.Encode2DArray(level.Levels))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("LevelAdded")
}

func (manager *Manager) RetrieveLevelByID(id string) (*model.Level, consts.ERRORMESSAGE) {
	manager.Init()
	defer manager.db.Close()

	//var level model.Level
	response, _ := manager.db.Query("select level from levels where id = ?", id)
	var levelString string
	for response.Next() {
		response.Scan(&levelString)
	}

	if levelString != "" {
		return &model.Level{ID: id, Levels: helper.Decode2DArray(levelString)}, ""
	}
	return nil, consts.LEVELNOTFOUND
}

func (manager *Manager) UpdateALevel(level *model.Level) consts.ERRORMESSAGE {
	manager.Init()
	defer manager.db.Close()

	//var level model.Level
	_, err := manager.db.Exec("update levels set level = ? where id = ?", helper.Encode2DArray(level.Levels), level.ID)
	if err != nil {
		return consts.LEVELNOTFOUND
	}
	return consts.LEVELUPDATED
}

func (manager *Manager) DeleteALevel(id string) consts.ERRORMESSAGE {
	manager.Init()
	defer manager.db.Close()

	//var level model.Level
	result, err := manager.db.Exec("DELETE FROM levels where id = ?", id)
	if err != nil {
		return consts.LEVELNOTFOUND
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		return consts.LEVELDELETED
	}
	return consts.LEVELNOTFOUND
}
