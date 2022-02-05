package alexa

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Request struct {
}

const (
	EventOnLaunch       = "OnLaunch"
	EventOnIntent       = "OnIntent"
	EventOnSessionEnded = "OnSessionEnded"
	EventOnAuthCheck    = "OnAuthCheck"
	AmazonCancelIntent  = "AMAZON.CancelIntent"
	AmazonHelpIntent    = "AMAZON.HelpIntent"
	AmazonNextIntent    = "AMAZON.NextIntent"
	AmazonStopIntent    = "AMAZON.StopIntent"
)

func handleSlateOne(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	c.JSON(
		http.StatusOK,
		gin.H{
			"version": "1.0",
			"response": gin.H{
				"shouldEndSession": true,
			},
		},
	)
}
