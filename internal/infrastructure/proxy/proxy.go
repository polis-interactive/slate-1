package proxy

import (
	grpcControl "github.com/polis-interactive/slate-1/api/v1/go"
	"sync"
)

type Proxy struct {
	mu               *sync.RWMutex
	connections      []*connection
	nextConnectionId uint32
}

func NewProxy() *Proxy {
	return &Proxy{
		mu:               &sync.RWMutex{},
		connections:      make([]*connection, 0),
		nextConnectionId: 0,
	}
}

func (p *Proxy) HandleConnectionOpen() (ch chan *grpcControl.ControlResponse, connectionId uint32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	newConnectionId := p.nextConnectionId
	p.nextConnectionId += 1
	newConn := newConnection(newConnectionId)
	p.connections = append(p.connections, newConn)
	return newConn.ch, newConnectionId
}

func (p *Proxy) HandleConnectionClose(connectionId uint32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	// nothing to do... shouldn't ever get here
	closePosition := -1
	for i, c := range p.connections {
		if c.connectionId == connectionId {
			closePosition = i
			c.Disconnect()
			break
		}
	}
	// again, shouldn't ever get here but
	if closePosition == -1 {
		return
	}
	p.connections = append(p.connections[:closePosition], p.connections[closePosition+1:]...)
}

func (p *Proxy) HandleDispatchMessage(response *grpcControl.ControlResponse) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, c := range p.connections {
		c.ch <- response
	}
}

func (p *Proxy) Shutdown() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, c := range p.connections {
		c.Disconnect()
	}
}
