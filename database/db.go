package database

import (
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=final_task_db port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Photo{})
	if err != nil {
		panic(err)
	}
}
