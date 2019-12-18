package main

import (
	_ "net/http/pprof"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile/v2"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/menta2l/lcm/lcm/cmd"
)

func main() {
	logr.SetLevelFromEnv()

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
