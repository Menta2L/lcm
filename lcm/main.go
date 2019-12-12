package main

import (
	_ "net/http/pprof"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile/v2"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/lcm/cmd"
	"github.com/menta2l/lcm/server"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.LcmServer{}

	lile.Name("lcm")
	lile.Server(func(g *grpc.Server) {
		lcm.RegisterLcmServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
