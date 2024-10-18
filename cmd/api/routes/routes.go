package routes

import (
	"goWallet/internal/handler"
	"goWallet/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	userService := services.Services{}

	userHandler := handler.Handler{
		User: userService,
	}

	api.GET("/status")
	api.POST("/register", userHandler.Register)
	return r
}
