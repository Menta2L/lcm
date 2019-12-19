package worker

import (
	"context"
	"fmt"

	"github.com/lileio/logr"
	"github.com/menta2l/lcm/pkg/models"
)

var JobQueue = make(chan CertReq, 10)

type CertReq struct {
	Cert     models.Cert
	Issuer   models.Issuer
	Solver   models.Solver
	Domains  []string
	Renewal  bool
	Attempts int
}

// NewCertReq constructs an instance of the CertReq struct
func NewCertReq(domains []string, renewal bool) CertReq {
	return CertReq{
		Domains:  domains,
		Renewal:  renewal,
		Attempts: 0,
	}
}

// Submit a job to the queue for processing
func (r *CertReq) Submit() {
	var operation string

	if r.Renewal {
		operation = "renewal"
	} else {
		operation = "initial certificates"
	}

	logr.WithCtx(context.Background()).Infof("Submitting '%v' to job queue for %s", r.Domains, operation)
	JobQueue <- *r
}

// CertWorker is the actual worker routine that eats away from the queue
func CertWorker(id int) {
	var prefix = fmt.Sprintf("[worker-%d]", id)
	for job := range JobQueue {
	}
}
