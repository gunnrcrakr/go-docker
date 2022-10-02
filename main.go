package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	dsn := "user:mypassword@tcp(db:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	// check mysql ping
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicln(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Connection Opened to Database ✅")

	// sleep for test only!
	time.Sleep(5 * time.Second)
}
