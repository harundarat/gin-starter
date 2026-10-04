package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postForm(c *gin.Context) {
	message := c.PostForm("message")
	nickname := c.DefaultPostForm("nickname", "anonymous")

	c.JSON(http.StatusCreated, gin.H{
		"message":  message,
		"nickname": nickname,
		"status":   "posted",
	})
}

func main() {
	r := gin.Default()

	r.POST("/form-post", postForm)

	r.Run(":8080")
}
