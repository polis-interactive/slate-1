package cloud

import "github.com/polis-interactive/slate-1/internal/types"

type Config struct {
	IpInterface string
	GrpcPort    int
	TLSConfig   *types.TLSConfig
}

func (c *Config) GetIpInterface() string {
	return c.IpInterface
}

func (c *Config) GetGrpcPort() int {
	return c.GrpcPort
}

func (c *Config) GetGrpcTLSConfig() *types.TLSConfig {
	return c.TLSConfig
}
