package routes

import (
	"goWallet/cmd/api/handler/user"
	postDB "goWallet/internal/repositories/postgres"
	"goWallet/internal/repositories/postgres/user"
	"goWallet/internal/services/user"

	"github.com/gin-gonic/gin"

	"log"
	"os"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	DSN := os.Getenv("DSN")

	client, err := postDB.ConnectDB(DSN)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := postgres.Repository{
		Client: client,
	}

	userService := services.Services{
		Repo: userRepository,
	}

	userHandler := handler.Handler{
		User: userService,
	}

	api.GET("/status")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	return r
}
