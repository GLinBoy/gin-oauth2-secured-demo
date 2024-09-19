package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-oauth2-secured-demo/config"
	"github.com/glinboy/gin-oauth2-secured-demo/dto"
	"github.com/glinboy/gin-oauth2-secured-demo/model"
)

var db = config.DatabaseConnecting()

func FindBooks(ctx *gin.Context) {
	var books []model.Book
	db.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(ctx *gin.Context) {
	var bookDTO dto.BookDTO
	if err := ctx.ShouldBindJSON(&bookDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := model.Book{Title: bookDTO.Title, Author: bookDTO.Author}
	db.Create(&book)

	ctx.JSON(http.StatusCreated, gin.H{"data": book})
}
