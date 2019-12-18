package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile/v2"
	"github.com/menta2l/lcm"
)

var s = LcmServer{}
var cli lcm.LcmClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		lcm.RegisterLcmServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = lcm.NewLcmClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
