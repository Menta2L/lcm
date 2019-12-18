package config

import (
	"fmt"
	"time"

	"github.com/lileio/lile"
)

type ClientConfig struct {
	Host string
	Port int
}

func (c *ClientConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type ServerConfig struct {
	lile.ServerConfig
	// How often to check for expired certificates
	ExpiryCheckAt time.Duration
	// How early before expiry shall certificates be renewed?
	RenewalThreshold time.Duration
}
type VaultConfig struct {
	Address string
	Token   string
}
