package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*------------ global database connection ------------*/
var DB *gorm.DB

func Connect() {
	/*------------ connect ------------*/
	/*------------ sqlite will create a new database if not existend ------------*/
	var err error
	DB, err = gorm.Open(sqlite.Open("discounts.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to database!")
		os.Exit(1)
	}

	/*------------ in case of updates, probably not needed ------------*/
	err = DB.AutoMigrate(
		&DiscountCode{},
	)
	if err != nil {
		fmt.Println("Database migration failed!")
		os.Exit(1)
	}
}
