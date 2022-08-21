package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type kp struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
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
func addKps(context *gin.Context) {
	var newKps kp

	if err := context.BindJSON(&newKps); err != nil {
		return
	}

	kps = append(kps, newKps)

	context.IndentedJSON(http.StatusCreated, newKps)
}

func main() {

	// router is the server
	router := gin.Default()

	router.GET("/kps", getKps)
	router.POST("/kps", addKps)

	// to run the server
	router.Run("localhost:5500")
}
