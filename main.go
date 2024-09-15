package main

import (
	"log"
	"time"

	"github.com/Dialosoft/src/app/config"
	"github.com/Dialosoft/src/app/database"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func main() {

	var err error
	var db *gorm.DB

	conf := config.GetNewDatabaseConfig()
	if conf.Database == "" {
		log.Fatal("not variable setted")
	}

	// Database
	for {
		var count int
		db, err = database.ConnectToDatabase(conf)
		if err == nil {
			break
		} else {
			count++
			time.Sleep(3 * time.Second)
		}
	}

	// Api Setup
	api := config.SetupAPI(db, fiber.Config{})

	if err := api.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
