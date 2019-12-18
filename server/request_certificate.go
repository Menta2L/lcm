package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/models"
)

func (s LcmServer) RequestCertificate(ctx context.Context, r *lcm.CertificateRequest) (*empty.Empty, error) {
	cert, err := models.NewCert(r, s.Backend)
	if err != nil {
		return nil, err
	}
	err = cert.Save(false)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
