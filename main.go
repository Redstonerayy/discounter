package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redstonerayy/discounter/db"
	"github.com/redstonerayy/discounter/routes"
)

func main() {
	/*------------ connect to database ------------*/
	db.Connect()

	/*------------ initialize gin and http routes ------------*/
	router := gin.Default()

	/*------------ static stuff ------------*/
	router.Static("/index.html", "./static/index.html")
	router.Static("/static", "./static")

	/*------------ api stuff ------------*/
	router.GET("/new-code", routes.NewCode)
	router.GET("/check-code", routes.CheckCode)

	/*------------ start gin ------------*/
	router.Run()
}
