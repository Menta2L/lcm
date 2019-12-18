package server

import (
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/vault"
)

type LcmServer struct {
	lcm.LcmServer
	Backend *vault.Client
}
