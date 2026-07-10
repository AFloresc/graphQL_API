package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("🚀 Starting Habit Streak Tracker GraphQL API...")

	// struct
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Habit Streak Tracker GraphQL API is running!",
			"status":  "success",
		})
	})

	router.GET("/statuss", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Habit Streak Tracker GraphQL API is running!",
			"status":  "success",
		})
	})

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
