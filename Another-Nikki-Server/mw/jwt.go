package mw

import (
	"Another-Nikki/util"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth(isRoot bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			util.SendResp(c, 404, nil, "You don't have permission")
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

		if !isRoot && mc.UserID > 2 {
			util.SendResp(c, 404, nil, "You don't have permission")
			c.Abort()
			return
		}
		articleID := c.Query("article_id")
		if !isRoot && articleID == "0" && mc.UserID > 2 {
			util.SendResp(c, 404, nil, "You don't have permission")
			c.Abort()
			return
		}

		c.Set("uid", mc.UserID)
		c.Set("username", mc.Username)

		AddVisitTime(mc.UserID, 1)
		c.Next()
	}
}
