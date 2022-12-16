package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/PavelDonchenko/3-event-micro/common"
	"github.com/gin-gonic/gin"
)

type tweetsSlice []common.Tweet

func (ts tweetsSlice) Len() int { //???????????????
	return len(ts)
}

func (ts tweetsSlice) Less(i, j int) bool {
	// After: for descending order
	return ts[i].CreatedAt.After(ts[j].CreatedAt)
}

func (ts tweetsSlice) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func sorting(data []common.Tweet) tweetsSlice {
	dateSortedReviews := make(tweetsSlice, 0, len(data))
	for _, d := range data {
		dateSortedReviews = append(dateSortedReviews, d)
	}
	sort.Sort(dateSortedReviews)
	return dateSortedReviews
}

func getDataFromRedisDatabase() []common.Tweet {
	var tweets []common.Tweet

	data := common.Get(common.InitRedisConn(), common.RedisDataKey)
	if len(data) == 0 {
		return tweets
	} else {
		err := json.Unmarshal([]byte(data), &tweets)
		common.HandleError(err, "GET REQUEST | Error occurred while unmarshal retrieved tweets from redis:database")
	}

	return tweets
}

func ginGetHttpRequest(context *gin.Context) {
	if len(getDataFromRedisDatabase()) == 0 {
		context.IndentedJSON(http.StatusInsufficientStorage, gin.H{"response": "redis database is empty"})
	} else {
		context.IndentedJSON(http.StatusOK, sorting(getDataFromRedisDatabase()))
	}
}

func getDataFromRedisDatabaseByParameter(context *gin.Context) {
	if len(getDataFromRedisDatabase()) == 0 {
		context.IndentedJSON(http.StatusInsufficientStorage, gin.H{"response": "redis database is empty"})
	} else {
		var founded []common.Tweet

		counter := 0
		for _, d := range getDataFromRedisDatabase() {
			if strings.EqualFold(d.Creator, context.Param("creator")) {
				founded = append(founded, d)
				counter++
			}
		}

		if counter == 0 {
			context.IndentedJSON(http.StatusNotFound, gin.H{"response": "record not found"})
		} else {
			context.IndentedJSON(http.StatusAccepted, sorting(founded))
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/tweet/list", ginGetHttpRequest)
	router.GET("/tweet/list/:creator", getDataFromRedisDatabaseByParameter)

	err := router.Run("localhost:9091")
	if err != nil {
		return
	}
}
