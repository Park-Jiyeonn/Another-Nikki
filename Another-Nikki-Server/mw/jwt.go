package mw

import (
	"Another-Nikki/util"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(404, gin.H{
				"message": "You don't have permission",
			})
			c.Abort()
			return
		}
		token := strings.Split(auth, " ")[1]

		mc, err := util.ParseToken(token)

		fmt.Println("mc =", mc)
		if err != nil {
			c.Abort()
			return
		}

		if mc.UserID > 2 {
			c.Abort()
			return
		}

		c.Set("uid", mc.UserID)
		c.Set("username", mc.Username)
		c.Next()
	}
}
