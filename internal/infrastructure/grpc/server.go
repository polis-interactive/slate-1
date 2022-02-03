package grpc

import (
	"errors"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcControl "github.com/polis-interactive/slate-1/api/v1/go"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type Server struct {
	grpcServer   *grpc.Server
	ip           string
	port         int
	shutdown     bool
	shutdownLock sync.Mutex
	shutdowns    chan struct{}
	wg           *sync.WaitGroup
}

func NewServer(
	config Config,
	proxy Proxy,
) (*Server, error) {

	log.Println("GrpcServer, NewServer: creating")

	var options []grpc.ServerOption
	options = append(
		options,
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcCtxTags.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxTags.UnaryServerInterceptor(),
		)),
	)

	shutdowns := make(chan struct{})
	wg := &sync.WaitGroup{}

	grpcSrv := grpc.NewServer(options...)

	controlService, err := newControlServer(proxy, shutdowns, wg)
	if err != nil {
		return nil, err
	}
	grpcControl.RegisterControlServer(grpcSrv, controlService)

	log.Println("GrpcServer, NewServer: created successfully")

	return &Server{
		grpcServer: grpcSrv,
		ip:         config.GetIpInterface(),
		port:       config.GetGrpcPort(),
		shutdown:   true,
		shutdowns:  shutdowns,
		wg:         wg,
	}, nil
}

func (s *Server) Startup() error {

	log.Println("GrpcServer, Startup: starting")

	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()

	if s.shutdown == false {
		return errors.New("GrpcServer, Startup: Tried to startup server twice")
	}

	addr := fmt.Sprintf("%s:%d", s.ip, s.port)
	log.Println(fmt.Sprintf("GrpcServer, Startup: listening at %s", addr))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("GrpcServer, Startup: Failed to listen: %v", err)
		return err
	}

	go func() {
		err := s.grpcServer.Serve(listener)
		if err != nil {
			log.Printf("GrpcServer: reported err %s", err)
			s.Shutdown()
		}
	}()

	s.shutdown = false

	log.Println("GrpcServer, Startup: started")

	return nil
}

func (s *Server) Shutdown() {

	log.Printf("GrpcServer, Shutdown: Shutting down")

	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()
	if s.shutdown {
		return
	}
	s.shutdown = true

	close(s.shutdowns)
	s.wg.Wait()
	s.grpcServer.Stop()

	log.Printf("GrpcServer, Shutdown: success")
}
