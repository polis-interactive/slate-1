package grpc

type Config interface {
	GetIpInterface() string
	GetGrpcPort() int
}
