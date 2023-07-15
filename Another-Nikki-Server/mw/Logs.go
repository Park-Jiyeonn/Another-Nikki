package mw

import (
	"Another-Nikki/dal"
	"github.com/gin-gonic/gin"
	"strings"
)

var log dal.Log

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		api := c.Request.Method + " " + c.Request.RequestURI
		platform := c.Request.Header.Get("Sec-Ch-Ua-Platform")
		platform = strings.Trim(platform, "\"")
		if platform == "" {
			platform = strings.Split(strings.Split(c.Request.Header.Get("User-Agent"), "(")[1], ";")[0]
		}
		_ = log.CreateLog(api, statusCode, clientIP, platform)
	}
}
