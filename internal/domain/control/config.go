package control

import (
	"github.com/polis-interactive/slate-1/internal/types"
)

type Config interface {
	GetGrpcServerAddress() string
	GetGrpcServerPort() int
	GetGrpcTLSConfig() *types.TLSConfig
}
