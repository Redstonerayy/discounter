package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redstonerayy/discounter/db"
)

/*------------ invalidate code if existend ------------*/
func InvalCode(c *gin.Context) {
	/*------------ parse client request ------------*/
	var codeinfo CodeInfo
	if err := c.BindJSON(&codeinfo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	/*------------ invalidate ------------*/
	status := db.Inval(codeinfo.Code)

	/*------------ return status of code ------------*/
	c.JSON(http.StatusOK, gin.H{
		"codestatus": status,
	})
}
