package alexa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func buildHandler(applicationId string) func(*gin.Context) {
	return func(c *gin.Context) {
		handleSlateOne(applicationId, c)
	}
}

type alexaApplication struct {
	id string `json:"applicationId" binding:"required"`
}

type alexaSession struct {
	application alexaApplication `json:"application" binding:"required"`
}

type lightIntentState struct {
	name  string `json:"name" binding:"required"`
	value string `json:"value" binding:"required"`
}

type lightIntentSlots struct {
	state lightIntentState `json:"state" binding:"required"`
}

type alexaIntentBody struct {
	name  string           `json:"name" binding:"required"`
	slots lightIntentSlots `json:"slots"`
}

type alexaRequest struct {
	requestType string          `json:"type" binding:"required"`
	intent      alexaIntentBody `json:"intent"`
}

type alexaBody struct {
	session alexaSession `json:"session" binding:"required"`
	request alexaRequest `json:"request" binding:"required"`
}

func handleSlateOne(applicationId string, c *gin.Context) {

	var body alexaBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.session.application.id != applicationId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unknown applicationId"})
		return
	}

	if body.request.requestType != "IntentRequest" {
		c.JSON(http.StatusOK, unhandledRequest())
		return
	}

	intentName := body.request.intent.name
	if intentName == "AMAZON.CancelIntent" || intentName == "AMAZON.StopIntent" {
		c.JSON(http.StatusOK, stopRequest())
		return
	} else if intentName == "AMAZON.HelpIntent" {
		c.JSON(http.StatusOK, askPolisRequest())
		return
	} else if intentName != "LightIntent" {
		c.JSON(http.StatusOK, unhandledRequest())
		return
	}

	newState := strings.ToLower(body.request.intent.slots.state.value)
	var isOff bool
	if newState == "on" || newState == "up" {
		isOff = false
	} else {
		isOff = true
	}

	log.Println(fmt.Sprintf("New slate state: %t", isOff))

	// try to dispatch to server
	if false {
		c.JSON(http.StatusOK, noSlateRequest())
		return
	}

	c.JSON(http.StatusOK, successRequest())
}

func unhandledRequest() map[string]interface{} {
	return gin.H{
		"version": "1.0",
		"response": gin.H{
			"shouldEndSession": true,
			"outputSpeech": gin.H{
				"type": "SSML",
				"ssml": "<speak><amazon:emotion name=\"excited\" intensity=\"low\">Sorry daddy, I don't know what to do</amazon:emotion></speak>",
			},
		},
	}
}

func stopRequest() map[string]interface{} {
	return gin.H{
		"version": "1.0",
		"response": gin.H{
			"shouldEndSession": true,
			"outputSpeech": gin.H{
				"type": "SSML",
				"ssml": "<speak><amazon:emotion name=\"disappointed\" intensity=\"low\">Sorry daddy, I'll stop</amazon:emotion></speak>",
			},
		},
	}
}

func askPolisRequest() map[string]interface{} {
	return gin.H{
		"version": "1.0",
		"response": gin.H{
			"shouldEndSession": true,
			"outputSpeech": gin.H{
				"type": "SSML",
				"ssml": "<speak><amazon:emotion name=\"excited\" intensity=\"high\">I know daddy, you should ask Polis!/amazon:emotion></speak>",
			},
		},
	}
}

func noSlateRequest() map[string]interface{} {
	return gin.H{
		"version": "1.0",
		"response": gin.H{
			"shouldEndSession": true,
			"outputSpeech": gin.H{
				"type": "SSML",
				"ssml": "<speak><amazon:emotion name=\"disappointed\" intensity=\"low\">Sorry daddy, I couldn't find slate one.../amazon:emotion></speak>",
			},
		},
	}
}

func successRequest() map[string]interface{} {
	return gin.H{
		"version": "1.0",
		"response": gin.H{
			"shouldEndSession": true,
			"outputSpeech": gin.H{
				"type": "SSML",
				"ssml": "<speak><amazon:emotion name=\"excited\" intensity=\"high\">Yes daddy!</amazon:emotion></speak>",
			},
		},
	}
}
