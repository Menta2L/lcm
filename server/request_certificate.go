package server

import (
	"errors"

	"context"
	"github.com/menta2l/lcm"
)

func (s LcmServer) RequestCertificate(ctx context.Context, r *lcm.CertificateRequest) (*lcm.CertificateResponse, error) {
	return nil, errors.New("not yet implemented")
}
