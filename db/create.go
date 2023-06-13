package db

import (
	"fmt"
	"os"

	"github.com/redstonerayy/discounter/codes"
)

func CreateDiscount() uint64 {
	/*------------ find a non-existent code ------------*/
	for {
		/*------------ new code ------------*/
		code := codes.GenCode()
		discountcode := DiscountCode{Code: code, Valid: true}

		/*------------ check if it exists ------------*/
		var count int64
		DB.Model(&DiscountCode{}).Where("Code = ?", code).Count(&count)

		/*------------ create code if it not exists ------------*/
		if !(count == 0) {
			res := DB.Create(&discountcode)
			if res.Error != nil {
				fmt.Println("Error creating discount code!")
				os.Exit(0)
			}

			return code
		}
	}
}
