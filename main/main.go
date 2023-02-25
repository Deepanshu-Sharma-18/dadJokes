package main

import (
	"github.com/Deepanshu-Sharma-18/api/database"
	"github.com/Deepanshu-Sharma-18/api/operations"
	"github.com/gin-gonic/gin"
)

func main() {
	database.StoreJokes()
	router := gin.Default()
	router.GET("/jokes", operations.GetJoke)
	router.GET("/random", operations.GetRandom)
	router.GET("/jokes/:id", operations.GetJokewithId)
	router.POST("/addJoke", operations.PostJoke)
	router.Run("localhost:8080")
}
