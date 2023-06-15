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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)

	/*------------ static stuff ------------*/
	router.Static("/app/", "./static")

	/*------------ api stuff ------------*/
	router.GET("/new-code", routes.NewCode)
	router.POST("/check-code", routes.CheckCode)
	router.POST("/inval-code", routes.InvalCode)
	router.POST("/reset-code", routes.ResetCode)

	/*------------ start gin ------------*/
	router.Run("127.0.0.1:8080")
}
