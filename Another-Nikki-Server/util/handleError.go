package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "404",
			"message": message +": " + err.Error(),
		})
		return true
	}
	return false
}