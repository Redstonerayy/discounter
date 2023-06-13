package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	/*------------ connect ------------*/
	var err error
	DB, err = gorm.Open(sqlite.Open("discounts.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to database!")
		os.Exit(1)
	}

	err = DB.AutoMigrate(
		&DiscountCode{},
	)
	if err != nil {
		fmt.Println("Database migration failed!")
		os.Exit(1)
	}
}
