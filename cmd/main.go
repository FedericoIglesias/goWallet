package main

import (
	"goWallet/cmd/api/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error: %s", err)
	}

	r := routes.SetupRouter()

	PORT := os.Getenv("PORT")

	r.Run(":" + PORT)

}
