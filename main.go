package main

import (
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/app"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/database"
	"log"
)

func main() {
	router := app.InitApp()

	database.InitDB()
	database.MigrateDB()

	log.Fatal(router.Run(":3000"))
}
