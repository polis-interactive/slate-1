package alexa

type Proxy interface {
	HandleAlexaCommand(isOn bool) error
}
