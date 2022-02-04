package grpc

import "github.com/polis-interactive/slate-1/internal/types"

type Config interface {
	GetIpInterface() string
	GetGrpcPort() int
	GetGrpcTLSConfig() *types.TLSConfig
}
