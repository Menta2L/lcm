package lcm

import (
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/lileio/lile/v2"
	"github.com/menta2l/lcm/pkg/config"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

var (
	cm     = &sync.Mutex{}
	Client LcmClient
)

func GetLcmClient(cfg config.ClientConfig) LcmClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
	}
	var serviceURL = ""
	if cfg.Host == "" {
		serviceURL = lile.URLForService("lcm")
	} else {
		serviceURL = cfg.Address()
	}

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, _ := grpc.Dial(
		serviceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				lile.ContextClientInterceptor(),
				otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
			),
		))

	cli := NewLcmClient(conn)
	Client = cli
	return cli
}
