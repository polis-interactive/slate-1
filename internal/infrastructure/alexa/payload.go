package alexa

type Application struct {
	Id string `json:"applicationId" binding:"required"`
}

type Session struct {
	Application Application `json:"application" binding:"required"`
}

type LightIntentState struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type LightIntentSlots struct {
	State LightIntentState `json:"state" binding:"required"`
}

type IntentBody struct {
	Name  string            `json:"name" binding:"required"`
	Slots *LightIntentSlots `json:"slots,omitempty"`
}

type Request struct {
	RequestType string      `json:"type" binding:"required"`
	Intent      *IntentBody `json:"intent,omitempty"`
}

type Body struct {
	Session Session `json:"session" binding:"required"`
	Request Request `json:"request" binding:"required"`
}
