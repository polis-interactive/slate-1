package control

type Config interface {
	GetGrpcServerAddress() string
}
