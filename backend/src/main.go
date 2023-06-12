package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/*------------ connect to database ------------*/

	/*------------ initialize gin and http routes ------------*/
	router := gin.Default()

	/*------------ static stuff ------------*/
	router.Static("/index.html", "./static/index.html")
	router.Static("/static", "./static")

	/*------------ api ------------*/
	router.GET("/new-code", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/check-code", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
