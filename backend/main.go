package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var r *gin.Engine

// var db *sql.DB

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:" + port) // Listen and Server at 8080/ping
}
