package routes

import (
	"goWallet/cmd/api/handler/account"
	"goWallet/cmd/api/handler/user"
	postDB "goWallet/internal/repositories/postgres"
	postgresAccount "goWallet/internal/repositories/postgres/account"
	"goWallet/internal/repositories/postgres/user"
	servicesAccount "goWallet/internal/services/account"
	"goWallet/internal/services/user"
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

	userRepository := postgres.Repository{
		Client: client,
	}

	userService := services.Services{
		Repo: userRepository,
	}

	userHandler := user.Handler{
		User: userService,
	}

	accountRepository := postgresAccount.Repository{
		Client: client,
	}

	accountService := servicesAccount.Services{
		Repo: accountRepository,
	}

	accountHandler := account.Handler{
		Account: accountService,
	}

	api.GET("/status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "STATUS UP"}) })
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/account", accountHandler.Create)
	return r
}
