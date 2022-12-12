package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/mayurkhairnar2525/restaurantManagement/modals"
	"log"
	"os"
)

// The purpose of this file is to return the variable db
// which help other files to interact with the database
var (
	db *gorm.DB
)

func InitDB() *gorm.DB {
	db = ConnectDB()
	return db
}

// ConnectDB will make connection with the database
func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: Unable to read the .env file variables", err)
	}

	// Get the value of an environment variable
	DriverName := os.Getenv("drivername")
	Root := os.Getenv("root")
	Password := os.Getenv("password")
	DbName := os.Getenv("databasename")

	db, err := gorm.Open(DriverName, Root+":"+Password+"@/"+DbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error", err)
	}

	db.AutoMigrate(
		&modals.User{})
	db.AutoMigrate(
		&modals.OrderApp{})

	fmt.Println("Database: Connected successfully")
	return db
}
