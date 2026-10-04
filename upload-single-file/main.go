package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("filename: %s", file.Filename)

	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		log.Printf("failed to save '%s' in the storage", file.Filename)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.String(http.StatusOK, "'%s' uploaded!", file.Filename)
}

func main() {
	router := gin.Default()

	router.POST("/upload", uploadFile)

	router.Run(":8080")
}
