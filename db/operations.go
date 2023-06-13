package db

import (
	"fmt"
	"os"

	"github.com/redstonerayy/discounter/codes"
)

func Create() uint64 {
	/*------------ find a non-existent code ------------*/
	for {
		/*------------ new code ------------*/
		code := codes.GenCode()
		discountcode := DiscountCode{Code: code, Valid: true}

		/*------------ check if it exists ------------*/
		var count int64
		DB.Model(&DiscountCode{}).Where("Code = ?", code).Count(&count)

		/*------------ use the code if it doesn't already exist ------------*/
		if count == 0 {
			/*------------ save in database ------------*/
			res := DB.Create(&discountcode)
			if res.Error != nil {
				fmt.Println("Error creating discount code!")
				os.Exit(0)
			}

			return code
		}
	}
}

func Check(code uint64) string {
	/*------------ check status of code ------------*/
	/*------------ check if it exists ------------*/
	var count int64
	DB.Model(&DiscountCode{}).Where("Code = ?", code).Count(&count)
	if count == 0 {
		return "not found"
	}

	/*------------ check if it's valid ------------*/
	var discountcode DiscountCode
	DB.Model(&DiscountCode{}).Where("Code = ?", code).First(&discountcode)

	if discountcode.Valid {
		return "valid"
	} else {
		return "invalid"
	}
}

func Inval(code uint64) string {
	/*------------ invalidate code ------------*/
	/*------------ check if it exists ------------*/
	var count int64
	DB.Model(&DiscountCode{}).Where("Code = ?", code).Count(&count)
	if count == 1 {
		/*------------ invalidate it ------------*/
		var discountcode DiscountCode
		DB.Model(&DiscountCode{}).Where("Code = ?", code).First(&discountcode)
		discountcode.Valid = false
		DB.Model(&DiscountCode{}).Where("Code = ?", code).Save(&discountcode)
		return "invalidated"
	} else {
		return "not found"
	}
}
