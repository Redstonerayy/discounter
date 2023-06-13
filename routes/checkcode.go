package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "check",
	})
}
