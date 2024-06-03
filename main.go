package main

import (
	"articles_pavlinov/database"
	"articles_pavlinov/routes"
	"os"
)

func main() {
	database.SetupDatabase()
	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
