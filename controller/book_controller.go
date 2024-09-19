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

func FindBook(ctx *gin.Context) {
	var book model.Book

	if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(ctx *gin.Context) {
	var book model.Book
	if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var bookDTO dto.BookDTO
	if err := ctx.ShouldBindJSON(&bookDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(bookDTO)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(ctx *gin.Context) {
	var book model.Book
	if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	ctx.JSON(http.StatusNoContent, gin.H{})
}
