package database

import (
	"log"
	"os"

	"github.com/sparsh459/REST-API-with-GO-FIber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// declare struct which will hold our database pointer
type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	// using gorm to open database and naming of db must end with {name of db}.db
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	// line to check error if connection is not happening
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	// print if no error coming
	log.Println("Connected Successfully to Database")
	// logger for our db, printing logs on command line and to know state of our database
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	// setting up our database
	// database that we get when we tell gorm to open
	Database = DbInstance{
		Db: db,
	}
}
