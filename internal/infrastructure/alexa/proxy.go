package alexa

type Proxy interface {
	HandleAlexaCommand(isOff bool) error
}
