package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type kp struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"status"`
}

var kps = []kp{
	{ID: "1", Item: "Blogs", Completed: false},
	{ID: "2", Item: "Video", Completed: false},
	{ID: "3", Item: "Sessions", Completed: false},
}

func getKps(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, kps)
}

func main() {

	// router is the server
	router := gin.Default()

	router.GET("/kps", getKps)

	// to run the server
	router.Run("localhost:5500")
}
