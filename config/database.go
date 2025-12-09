package config

import (
	"fmt"
	"go-package/model"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?parseTime=true",
		ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_DATABASE)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	DB = db

	log.Println("Database connected")

	// Lakukan
	autoMigrate(db)
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		// tambahkan model lainnya disini
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	log.Println("Database migration completed")
}
