package main

import (
	"fmt"
	"gaphql-api/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

// function stubs
// test

func main() {
	fmt.Println("🚀 Starting Habit Streak Tracker GraphQL API...")

	db, err := database.InitDB("./data/habits.db")
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	defer db.Close()
	fmt.Println("🚀 Databse connected successgully!")

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
