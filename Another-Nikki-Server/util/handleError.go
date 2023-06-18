package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"state":   "error",
			"message": message +": " + err.Error(),
		})
		return true
	}
	return false
}