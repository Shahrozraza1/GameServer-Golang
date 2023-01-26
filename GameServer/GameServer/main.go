package main

import (
	"GameServer/api"
	"GameServer/db"
)

func main() {
	db.GetManagerInstance().CreateTables()
	api.StartApiServer("127.0.0.1", "8080")
}
