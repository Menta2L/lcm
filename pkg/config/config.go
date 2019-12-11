package config

import (
	"strconv"
	"time"
)

// ServerConfig holds configuration that is required for running a server instance.
type ServerConfig struct {
	// Which port the server shall run on.
	Port int
	// How often to check for expired certificates
	ExpiryCheckAt time.Duration
	// How early before expiry shall certificates be renewed?
	RenewalThreshold time.Duration
}

// ClientConfig holds configuration that is required for creating a client
type ClientConfig struct {
	// Which host shall the client connect to?
	Hostname string
	// Which port is the host listening on?
	Port int
	// How often to check for expired certificates
	ExpiryCheckAt time.Duration
	// How early before expiry shall certificates be renewed?
	RenewalThreshold time.Duration
}

// GetDialAddr gets the formatted address to dial a new gRPC connection
func (c *ClientConfig) GetDialAddr() string {
	return c.Hostname + ":" + strconv.Itoa(c.Port)
}
