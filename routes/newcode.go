package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redstonerayy/discounter/db"
)

/*------------ create new code ------------*/
func NewCode(c *gin.Context) {
	/*------------ create ------------*/
	code := db.Create()

	/*------------ return code to client ------------*/
	c.JSON(http.StatusOK, gin.H{
		"newcode": code,
	})
}
