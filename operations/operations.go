package operations

import (
	"math/rand"
	"net/http"

	"github.com/Deepanshu-Sharma-18/api/model"
	"github.com/Deepanshu-Sharma-18/api/utils"
	"github.com/gin-gonic/gin"
)

func GetJoke(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, utils.DadJokes)
}

func GetRandom(ctx *gin.Context) {
	randnum := rand.Intn(len(utils.DadJokes) - 1)
	ctx.IndentedJSON(http.StatusOK, utils.DadJokes[randnum])
}

func GetJokewithId(ctx *gin.Context) {
	for i, t := range utils.DadJokes {
		if t.Id == ctx.Param("id") {
			ctx.IndentedJSON(http.StatusOK, utils.DadJokes[i])
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
	return

}

func PostJoke(ctx *gin.Context) {

	var newJokes model.Jokes
	newJokes.Id = string(utils.TotalIds)
	utils.TotalIds = utils.TotalIds + 1
	err := ctx.BindJSON(&newJokes)
	if err != nil {
		panic(err)
	}
	utils.DadJokes = append(utils.DadJokes, newJokes)
	ctx.IndentedJSON(http.StatusOK, newJokes)

}
