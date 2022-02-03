package proxy

import (
	grpcControl "github.com/polis-interactive/slate-1/api/v1/go"
	"sync"
)

type connection struct {
	connectionId uint32
	ch           chan *grpcControl.ControlResponse
	mu           *sync.RWMutex
}

func newConnection(connectionId uint32) *connection {
	conn := &connection{
		connectionId: connectionId,
		ch:           make(chan *grpcControl.ControlResponse, 10),
		mu:           &sync.RWMutex{},
	}
	return conn
}

func (c *connection) Disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ch != nil {
		close(c.ch)
		c.ch = nil
	}
}
