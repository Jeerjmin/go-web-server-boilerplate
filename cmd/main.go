package main

import (
	"go-auth/internal/config"
	database "go-auth/internal/db"
	"go-auth/internal/server"
)

func main() {
	c := config.NewConfig()

	db := database.NewDatabase(c)
	db.Connect()

	server.NewServer(c, db).Start()
}
