package mw

import (
	"Another-Nikki/dal"
	"bytes"

	"github.com/gin-gonic/gin"
)

var log dal.Log

type responseBodyWriter struct {
	gin.ResponseWriter  // 继承原有 gin.ResponseWriter
	bodyBuf *bytes.Buffer  // Body 内容临时存储位置，这里指针，原因这个存储对象要复用
}

// 覆盖原有 gin.ResponseWriter 中的 Write 方法
func (w *responseBodyWriter) Write(b []byte) (int, error) {
	if count, err := w.bodyBuf.Write(b); err != nil {  // 写入数据时，也写入一份数据到缓存中
		return count, err
	}
	return w.ResponseWriter.Write(b) // 原始框架数据写入
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := responseBodyWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = &writer

		c.Next()

		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		api := c.Request.Proto + " " + c.Request.Method + " " + c.Request.RequestURI
		
		log.CreateLog(api,statusCode, clientIP,writer.bodyBuf.String())
	}
}
