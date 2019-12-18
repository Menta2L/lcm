package server

import (
	"testing"

	"context"
	"github.com/menta2l/lcm"
	"github.com/stretchr/testify/assert"
)

func TestRequestCertificate(t *testing.T) {
	ctx := context.Background()
	req := &lcm.CertificateRequest{}

	res, err := cli.RequestCertificate(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
