package cloud

type Config struct {
	IpInterface string
	GrpcPort    int
}

func (c *Config) GetIpInterface() string {
	return c.IpInterface
}

func (c *Config) GetGrpcPort() int {
	return c.GrpcPort
}
