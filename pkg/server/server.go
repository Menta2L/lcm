package server

import (
	"github.com/menta2l/lcm/pkg/api"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/xenolf/lego/log"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

// StartServer spawns a server instance given a server config
func StartServer(config *config.ServerConfig, userAgent string) {
	s := grpc.NewServer()
	issuerService := NewIssuerService(config)
	api.RegisterIssuerServiceServer(s, issuerService)
	ln, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(config.Port))

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
