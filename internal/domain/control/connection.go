package control

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	grpcControl "github.com/polis-interactive/slate-1/api/v1/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"sync"
	"time"
)

type connection struct {
	serverAddress string
	bus           Bus
	mu            *sync.RWMutex
	wg            *sync.WaitGroup
	shutdowns     chan struct{}
}

func newConnection(conf Config, bus Bus) *connection {
	return &connection{
		serverAddress: conf.GetGrpcServerAddress(),
		bus:           bus,
		mu:            &sync.RWMutex{},
		wg:            &sync.WaitGroup{},
		shutdowns:     nil,
	}
}

func (c *connection) startup() {

	log.Println("Connection, startup; starting")

	if c.shutdowns == nil {
		c.shutdowns = make(chan struct{})
		c.wg.Add(1)
		go c.runMainLoop()
	}

	log.Println("Connection, startup; started")
}

func (c *connection) runMainLoop() {

	log.Println("ControlConnection, runMainLoop: running")

	for {
		func(c *connection) {
			conn, err := c.tryConnectToServer()
			if err != nil {
				return
			}
			defer c.tryCloseConnection(conn)
			c.streamClient(conn)
		}(c)
		select {
		case _, ok := <-c.shutdowns:
			if !ok {
				goto finish
			}
		case <-time.After(5 * time.Second):
			log.Println("ControlConnection, runMainLoop: retrying connection")
		}
	}
finish:
	c.wg.Done()
}

func (c *connection) shutdown() {

	log.Println("Connection, shutdown; shutting down")

	if c.shutdowns != nil {
		close(c.shutdowns)
		c.wg.Wait()
		c.shutdowns = nil
	}
	log.Println("Connection, shutdown; finished")
}

func (c *connection) tryConnectToServer() (*grpc.ClientConn, error) {

	log.Println("ControlConnection, tryConnectToServer: connecting")

	var opts []grpc.DialOption
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(c.serverAddress, opts...)
	if err != nil {
		log.Println(fmt.Sprintf("ControlConnection, tryConnectToServer: error while dialing; %s", err.Error()))
		return nil, err
	}

	log.Println("ControlConnection, tryConnectToServer: connected")

	return conn, nil

}

func (c *connection) tryCloseConnection(conn *grpc.ClientConn) {

	log.Println("ControlConnection, tryCloseConnection: closing")

	err := conn.Close()
	if err != nil {
		log.Println(fmt.Sprintf("ControlConnection, tryCloseConnection: error while closing client; %s", err.Error()))
	}
}

func (c *connection) streamClient(conn *grpc.ClientConn) {

	log.Println("ControlConnection, streamClient: setting up")

	client := grpcControl.NewControlClient(conn)
	stream, err := client.ControlConnection(context.Background(), &grpcControl.EmptyRequest{})
	if err != nil {
		log.Println(fmt.Sprintf("ControlConnection, streamClient: error while connecting; %s", err.Error()))
		return
	}

	commandsIn := make(chan *grpcControl.ControlResponse, 10)
	c.wg.Add(1)

	go func(
		stream grpcControl.Control_ControlConnectionClient,
		ch chan *grpcControl.ControlResponse,
		wg *sync.WaitGroup,
	) {
		for {
			response, e := stream.Recv()
			if e == io.EOF {
				goto StopStream
			}
			if e != nil {
				log.Println(fmt.Sprintf("ControlConnection, streamClient, in chan: reported err %s", e.Error()))
				goto StopStream
			}
			ch <- response
		}
	StopStream:
		log.Println("ControlConnection, streamClient, in chan: closing")
		close(ch)
		wg.Done()
	}(stream, commandsIn, c.wg)

	log.Println("ControlConnection, streamClient: handling messages")

	err = func(
		stream grpcControl.Control_ControlConnectionClient,
		bus Bus,
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
			case command, ok := <-commandsIn:
				if !ok {
					return errors.New("command in channel unexpectedly closed")
				}
				bus.HandleControlResponse(command)
			}
		}
	}(stream, c.bus, c.shutdowns)

	if err != nil {
		log.Println(fmt.Sprintf("ControlConnection, streamClient: reported err; %s", err.Error()))
	}
}
