package alexa

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func buildHandler(applicationId string) func(*gin.Context) {
	return func(c *gin.Context) {
		handleSlateOne(applicationId, c)
	}
}

func handleSlateOne(applicationId string, c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"version": "1.0",
			"response": gin.H{
				"shouldEndSession": true,
				"outputSpeech": gin.H{
					"type": "SSML",
					"ssml": "<speak><amazon:emotion name=\"excited\" intensity=\"high\">Yes daddy!</amazon:emotion></speak>",
				},
			},
		},
	)
}
