package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(ctx *gin.Context) {
		ids := ctx.QueryMap("ids")
		names := ctx.PostFormMap("names")

		ctx.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	r.Run(":8080")
}
