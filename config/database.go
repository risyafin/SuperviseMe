package config

import (
	"fmt"
	"log"
	"os"
	"superviseMe/core/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbName     = "DB_NAME"
	ServerPort = "DB_PORT"
)

func GetConfig() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", os.Getenv(dbUsername), os.Getenv(dbPassword),
		os.Getenv(dbHost), os.Getenv(ServerPort), os.Getenv(dbName))
}

func AutoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Notification{},
		&entity.TypeNotification{},
		&entity.ActivityLog{},
		&entity.Goals{},
		&entity.List{},
		&entity.Comment{},
		&entity.Card{},
		&entity.CheckListCard{},
	)
	if err != nil {
		log.Fatalf("Migration Failed. error: %v", err)
	}
	log.Println("Migration Succes...")
}

func InitDatabaseConnection(string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(GetConfig()), &gorm.Config{})
	if err != nil {
		panic("Connection Failed")
	}
	return DB
}
