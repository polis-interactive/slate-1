package alexa

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type handler struct {
	applicationId string
}

func (h *handler) handleSlateOne(c *gin.Context) {

	b, _ := ioutil.ReadAll(c.Request.Body)
	println(string(b))
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

	var body alexaBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(fmt.Sprintf("Failed to decode json with err: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(body)

	appId := body.session.application.id
	if appId != h.applicationId {
		log.Println(fmt.Sprintf("Unknown application id: %s", appId))
		c.JSON(http.StatusForbidden, gin.H{"error": "unknown applicationId"})
		return
	}

	requestType := body.request.requestType
	if requestType != "IntentRequest" {
		log.Println(fmt.Sprintf("Unknown request type: %s", requestType))
		c.JSON(http.StatusOK, unhandledResponse())
		return
	}

	intentName := body.request.intent.name
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
		c.JSON(http.StatusOK, noSlateResponse())
		return
	}

	c.JSON(http.StatusOK, successResponse())
}
