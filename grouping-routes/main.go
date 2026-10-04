package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// login, submit, read
func loginEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "login"})
}

func submitEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "submit"})
}

func readEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"action": "read"})
}

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("ini telah melewati auth required middleware")
		ctx.Next()
	}
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/login", loginEndpoint)
	v1.POST("/submit", submitEndpoint)
	v1.POST("/read", readEndpoint)

	v2 := router.Group("v2")
	v2.Use(AuthRequired())

	v2.POST("/login", loginEndpoint)
	v2.POST("/submit", submitEndpoint)
	v2.POST("/read", readEndpoint)

	router.Run(":8080")
}
