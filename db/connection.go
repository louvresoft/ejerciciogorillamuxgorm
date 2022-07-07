package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DBS = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5433"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DBS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Conectado a DB..")
	}
}
