package cloud

import "github.com/polis-interactive/slate-1/internal/types"

type Config struct {
	IpInterface   string
	GrpcPort      int
	TLSConfig     *types.TLSConfig
	AlexaPort     int
	IsProduction  bool
	ApplicationId string
}

func (c *Config) GetAlexaPort() int {
	return c.AlexaPort
}

func (c *Config) GetIsProduction() bool {
	return c.IsProduction
}

func (c *Config) GetApplicationId() string {
	return c.ApplicationId
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
