package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/test", func(cx *gin.Context){
		cx.JSON(200, gin.H{
			"message":"High Fidelity music!",
		})
	})

	server.Run(":8000")
}