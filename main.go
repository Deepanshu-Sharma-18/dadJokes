package main

import (
	"os"

	"github.com/Deepanshu-Sharma-18/api/database"
	"github.com/Deepanshu-Sharma-18/api/operations"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.StoreJokes()
	router := gin.Default()
	router.GET("/jokes", operations.GetJoke)
	router.GET("/random", operations.GetRandom)
	router.GET("/jokes/:id", operations.GetJokewithId)
	router.POST("/addJoke", operations.PostJoke)
	port := os.Getenv("PORT")
	router.Run("localhost:%v", port)
}
