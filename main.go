package main

import (
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/app"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/database"
	"log"
)

func main() {
	server := app.InitApp()

	database.InitDB()
	database.MigrateDB()

	println("Server is running on port 3000")

	log.Fatal(server.Run(":3000"))
}
