package alexa

type Config interface {
	GetAlexaPort() int
	GetIsProduction() bool
	GetApplicationId() string
	GetIpInterface() string
}
