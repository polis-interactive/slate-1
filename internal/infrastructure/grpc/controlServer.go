package grpc

import (
	"fmt"
	grpcControl "github.com/polis-interactive/slate-1/api/v1/go"
	"github.com/polis-interactive/slate-1/internal/infrastructure"
	"log"
	"sync"
)

type controlServer struct {
	grpcControl.UnsafeControlServer
	p         Proxy
	shutdowns chan struct{}
	wg        *sync.WaitGroup
}

var _ grpcControl.ControlServer = (*controlServer)(nil)

func newControlServer(proxy Proxy, shutdowns chan struct{}, wg *sync.WaitGroup) (*controlServer, error) {

	log.Println("participantServer, newServer: creating")

	return &controlServer{
		p:         proxy,
		shutdowns: shutdowns,
		wg:        wg,
	}, nil
}

func (c *controlServer) ControlConnection(_ *grpcControl.EmptyRequest, stream grpcControl.Control_ControlConnectionServer) error {

	log.Println("controlServer, ControlConnection: starting")

	controlChannel, connectionId := c.p.HandleConnectionOpen()

	log.Println(fmt.Sprintf("controlServer, ControlConnection: new connection with id %d", connectionId))

	c.wg.Add(1)

	err := func(
		stream grpcControl.Control_ControlConnectionServer,
		ch chan *grpcControl.ControlResponse,
		chId uint32,
		shutdowns chan struct{},
	) error {
		for {
			select {
			case _, ok := <-shutdowns:
				if !ok {
					return nil
				}
			case <-stream.Context().Done():
				return nil
			case msg, ok := <-ch:
				if !ok {
					return &infrastructure.ErrorChannelClosed{
						ChannelId: connectionId,
					}
				}
				err := stream.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}(stream, controlChannel, connectionId, c.shutdowns)

	if err != nil {
		log.Println(err)
	}

	log.Println(fmt.Sprintf("controlServer, ControlConnection: closing connection with id %d", connectionId))

	c.p.HandleConnectionClose(connectionId)

	c.wg.Done()

	return err
}
