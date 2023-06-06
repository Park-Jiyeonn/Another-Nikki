package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func GithubWebhook(c *gin.Context) {
	header := c.Request.Header
	fmt.Println("=================================================")
	fmt.Println(header)
	b, _ := io.ReadAll(c.Request.Body)
	fmt.Println(string(b))
	fmt.Println("=================================================")
}
