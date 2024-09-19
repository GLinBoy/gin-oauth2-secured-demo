package router

import (
	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-oauth2-secured-demo/controller"
)

func NewRouter() *gin.Engine {
	routes := gin.Default()

	apiRoutes := routes.Group("/api")

	bookRoutes := apiRoutes.Group("/books")
	bookRoutes.GET("", controller.FindBooks)
	bookRoutes.POST("", controller.CreateBook)
	bookRoutes.GET("/:id", controller.FindBook)

	return routes
}
