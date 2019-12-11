package config

import "time"

// ServerConfig holds configuration that is required for running a server instance.
type ServerConfig struct {
	// Which port the server shall run on.
	Port int
	// How often to check for expired certificates
	ExpiryCheckAt time.Duration
	// How early before expiry shall certificates be renewed?
	RenewalThreshold time.Duration
}
