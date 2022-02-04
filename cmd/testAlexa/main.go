package main

import alexa "github.com/mikeflynn/go-alexa/skillserver"

var Applications = map[string]interface{}{
	"/slate-1/LightIntent": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.69a5128a-d6b6-4bd2-888d-f388e8986c7b", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}

func main() {
	alexa.Run(Applications, "420")
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}
