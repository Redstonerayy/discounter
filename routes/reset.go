package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redstonerayy/discounter/db"
)

/*------------ invalidate code if existend ------------*/
func ResetCode(c *gin.Context) {
	/*------------ parse client request ------------*/
	var codeinfo CodeInfo
	if err := c.BindJSON(&codeinfo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	/*------------ invalidate ------------*/
	status := db.Reset(codeinfo.Code)

	/*------------ return status of code ------------*/
	c.JSON(http.StatusOK, gin.H{
		"codestatus": status,
	})
}
