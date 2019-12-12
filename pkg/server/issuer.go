package server

import (
	"context"

	"github.com/menta2l/lcm/pkg/api"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/mholt/certmagic"
	"github.com/xenolf/lego/log"
)

// IssuerService issues certificates to clients
type IssuerService struct {
	client *certmagic.Config
}

// NewIssuerService constructs a new instance with a predefined config
func NewIssuerService(config *config.ServerConfig) *IssuerService {
	issuer := new(IssuerService)
	return issuer

}
func (s *IssuerService) CreateIssuer(ctx context.Context, in *api.IssuerRequest) (r *api.IssuerResponse, err error) {
	if in.GetSelfSignedIssuer() != nil {
		log.Infof("Creating new self signed issuer %s", in.GetSelfSignedIssuer().GetName())
	} else {
		log.Infof("Creating acme issuer")
	}
	return &api.IssuerResponse{}, nil

}
func (s *IssuerService) ListIssuers(context.Context, *api.ListIssuerRequest) (r *api.ListIssuerCollection, err error) {
	return
}

func (s *IssuerService) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
	log.Infof("Ping request")
	return &api.PingResponse{Msg: "pong"}, nil
}
