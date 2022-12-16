package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PavelDonchenko/3-event-micro/common"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/tweet", postHTTPRequest)

	err := router.Run(":9090")
	if err != nil {
		log.Fatal("Error create route")
	}
}

func postHTTPRequest(c *gin.Context) {
	tweet := common.Tweet{}

	err := c.BindJSON(&tweet)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"response": "bad request"})
	}

	tNow := time.Now().Format("2006-01-02 15:04:05")
	tweet.CreatedAt, err = time.Parse("2006-01-02 15:04:05", tNow)
	common.HandleError(err, "Error occurred while parsing the creation time")

	c.IndentedJSON(http.StatusCreated, tweet)

	common.Producer(common.AmqpUrl, common.RabbitQueueName, tweet)
}
