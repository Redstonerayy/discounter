package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redstonerayy/discounter/db"
)

func NewCode(c *gin.Context) {
	code := db.CreateDiscount()

	fmt.Println(code)

	c.JSON(http.StatusOK, gin.H{
		"message": "new",
		"code":    code,
	})
}
