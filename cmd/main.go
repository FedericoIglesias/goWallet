package main

import (
	"goWallet/internal/db"
	"goWallet/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error: %s", err)
	}

	db.ConnectDB()

	r := routes.SetupRouter()

	PORT := os.Getenv("PORT")

	r.Run(":" + PORT)

}
