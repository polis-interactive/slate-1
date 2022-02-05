package alexa

import "github.com/gin-gonic/gin"

func unhandledResponse() map[string]interface{} {
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

func stopResponse() map[string]interface{} {
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

func askPolisResponse() map[string]interface{} {
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

func noSlateResponse() map[string]interface{} {
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

func successResponse() map[string]interface{} {
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
