package main

import "goWallet/internal/services"

// "goWallet/cmd/api/routes"
// "log"
// "os"

// "github.com/joho/godotenv"

func main() {

	str := services.CreateAlias()

	println(str)

	// if err := godotenv.Load("./.env"); err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }

	// r := routes.SetupRouter()

	// PORT := os.Getenv("PORT")

	// r.Run(":" + PORT)

}
