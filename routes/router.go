package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaiminbhaduri/codeprac/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Views
	router.GET("/", controllers.Dashboard)

	// API
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	router.POST("/api/execute", controllers.ExecuteCode)

	return router
}
