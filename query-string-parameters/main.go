package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")

		ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.Run(":8080")
}
