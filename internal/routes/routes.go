package routes

import (
	"goWallet/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/status")
	api.POST("/register",controllers.Register)
	return r
}
