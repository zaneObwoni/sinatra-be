package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//run database
	configs.ConnectDB()

	router := gin.Default()

	router.Run("localhost:8000")
}
