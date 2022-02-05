package alexa

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
	slots lightIntentSlots `json:"slots,omitempty"`
}

type alexaRequest struct {
	requestType string          `json:"type" binding:"required"`
	intent      alexaIntentBody `json:"intent,omitempty"`
}

type alexaBody struct {
	version string       `json:"version"`
	session alexaSession `json:"session" binding:"required"`
	request alexaRequest `json:"request" binding:"required"`
}
