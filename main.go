package main

import (
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/app"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/database"
	"log"
)

func main() {
	server := app.InitApp()

	database.InitDB()
	database.MigrateDB()

	println("Server is running on port 3000")

	log.Fatal(server.Run(":3000"))
}
