package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func getUser(ctx *gin.Context) {
	name := ctx.Param("name")
	greeting := fmt.Sprintf("Hello %s", name)
	ctx.JSON(http.StatusOK, gin.H{"message": greeting})
}

func getUserAction(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")
	userAction := fmt.Sprintf("User %s is %s", name, action)

	ctx.String(http.StatusOK, userAction)
}

func posting(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"method": "POST"})
}

func putting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func deleting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func patching(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}

func head(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func options(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/some-get", getting)
	r.GET("/user/:name", getUser)
	r.GET("/user/:name/*action", getUserAction)
	r.POST("/some-post", posting)
	r.PUT("/some-put", putting)
	r.DELETE("/some-delete", deleting)
	r.PATCH("/some-patch", patching)

	r.HEAD("/some-head", head)
	r.OPTIONS("/some-options", options)

	r.Run()
}
