package client

import (
	"context"
	"errors"
	"github.com/menta2l/lcm/pkg/api"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/xenolf/lego/log"
	"google.golang.org/grpc"
	"os"
	"time"
)

// StartClient starts a client instance with a client config and a user agent
func StartClient(config *config.ClientConfig, userAgent string) {
	log.Infof("Initializing client")

	// Check if the config is valid
	err := validateConfig(config)

	if err != nil {
		log.Warnf("Configuration error: %v", err)
		os.Exit(1)
	}

	log.Infof("Configuring connection to %s for gRPC operations", config.GetDialAddr())

	// Configure connection
	conn, err := grpc.Dial(config.GetDialAddr(), grpc.WithInsecure(), grpc.WithUserAgent(userAgent+";")) // TODO: Not run insecure

	if err != nil {
		log.Warnf("Could not configure connection to host: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	// Create client from connection
	client := api.NewIssuerServiceClient(conn)

	// Test connection
	_, err = client.Ping(context.Background(), &api.PingRequest{Msg: "ping"})

	if err != nil {
		log.Warnf("Could not test connection to %s: %v", config.GetDialAddr(), err)
		os.Exit(1)
	}

	log.Infof("Successfully pinged lcm server")
}

func validateConfig(c *config.ClientConfig) error {
	if len(c.Hostname) == 0 {
		return errors.New("hostname was empty")
	}

	if c.RenewalThreshold > (24*time.Hour)*30 {
		return errors.New("renewal threshold can't exceed 30 days")
	}

	return nil
}
