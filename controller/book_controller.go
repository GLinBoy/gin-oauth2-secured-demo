package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-oauth2-secured-demo/config"
	"github.com/glinboy/gin-oauth2-secured-demo/model"
)

var db = config.DatabaseConnecting()

func FindBooks(ctx *gin.Context) {
	var books []model.Book
	db.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
