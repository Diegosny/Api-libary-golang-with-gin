package main

import (
	"api/database"
	"api/server"
)

func main() {
	database.StartDB()

	newServer := server.NewServer()
	newServer.Run()
}
