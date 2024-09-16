package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	routes := gin.Default()

	return routes
}
