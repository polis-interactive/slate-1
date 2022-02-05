package alexa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type handler struct {
	applicationId string
}

func (h *handler) handleSlateOne(c *gin.Context) {

	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(fmt.Sprintf("Failed to decode json with err: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(body)

	appId := body.Session.Application.Id
	if appId != h.applicationId {
		log.Println(fmt.Sprintf("Unknown application id: %s", appId))
		c.JSON(http.StatusForbidden, gin.H{"error": "unknown applicationId"})
		return
	}

	requestType := body.Request.RequestType
	if requestType != "IntentRequest" {
		log.Println(fmt.Sprintf("Unknown request type: %s", requestType))
		c.JSON(http.StatusOK, unhandledResponse())
		return
	}

	intentName := body.Request.Intent.Name
	if intentName == "AMAZON.CancelIntent" || intentName == "AMAZON.StopIntent" {
		log.Println(fmt.Sprintf("Stop intent requested"))
		c.JSON(http.StatusOK, stopResponse())
		return
	} else if intentName == "AMAZON.HelpIntent" {
		log.Println(fmt.Sprintf("Help intent requested"))
		c.JSON(http.StatusOK, askPolisResponse())
		return
	} else if intentName != "LightIntent" {
		log.Println(fmt.Sprintf("unknown intent requested"))
		c.JSON(http.StatusOK, unhandledResponse())
		return
	}

	newState := strings.ToLower(body.Request.Intent.Slots.State.Value)
	var isOff bool
	if newState == "on" || newState == "up" {
		isOff = false
	} else {
		isOff = true
	}

	log.Println(fmt.Sprintf("New slate state: %t", isOff))

	// try to dispatch to server
	if false {
		c.JSON(http.StatusOK, noSlateResponse())
		return
	}

	c.JSON(http.StatusOK, successResponse())
}
