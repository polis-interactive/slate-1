package alexa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type handler struct {
	appId string
	p     Proxy
}

func (h *handler) handleSlateOne(c *gin.Context) {

	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(fmt.Sprintf("Failed to decode json with err: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appId := body.Session.Application.Id
	if appId != h.appId {
		log.Println(fmt.Sprintf("Unknown application id: %s", appId))
		c.JSON(http.StatusForbidden, gin.H{"error": "unknown applicationId"})
		return
	}

	requestType := body.Request.RequestType
	if requestType == "SessionEndedRequest" {
		log.Println(fmt.Sprintf("SessionEnded"))
		return
	} else if requestType == "LaunchRequest" {
		log.Println("Launch request")
		c.JSON(http.StatusOK, launchResponse())
		return
	} else if requestType != "IntentRequest" {
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
	var isOn bool
	if newState == "on" || newState == "up" {
		isOn = true
	} else {
		isOn = false
	}

	err := h.p.HandleAlexaCommand(isOn)
	if err != nil {
		log.Println(fmt.Sprintf("No slate found"))
		c.JSON(http.StatusOK, noSlateResponse())
		return
	}

	c.JSON(http.StatusOK, successResponse())
}
