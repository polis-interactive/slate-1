package grpc

import (
	grpcControl "github.com/polis-interactive/slate-1/api/v1/go"
)

type Proxy interface {
	HandleConnectionOpen() (ch chan *grpcControl.ControlResponse, connectionId uint32)
	HandleConnectionClose(connectionId uint32)
}
