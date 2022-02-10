package main

import (
	"log"
	"os"
	"sinatra/configs"
	"sinatra/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//run database
	configs.ConnectDB()
	log.Print(os.Getenv("MONGOURI"))

	router := gin.Default()

	//routes
	routes.UserRoute(router)

	router.Run("localhost:8080")
}
