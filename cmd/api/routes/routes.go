package routes

import (
	handlerUser "goWallet/cmd/api/handler/user"
	postDB "goWallet/internal/repositories/postgres"
	postgresAccount "goWallet/internal/repositories/postgres/account"
	postgresUser "goWallet/internal/repositories/postgres/user"
	
	servicesUser "goWallet/internal/services/user"
	"net/http"

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

	userRepository := postgresUser.Repository{
		Client: client,
	}

	accountRepository := postgresAccount.Repository{
		Client: client,
	}

	userService := servicesUser.Services{
		Repo: userRepository,
	}

	// accountService := servicesAccount.Services{
	// 	Repo: accountRepository,
	// }

	userHandler := handlerUser.Handler{
		User:    userService,
	}

	api.GET("/status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "STATUS UP"}) })
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)

	return r
}
