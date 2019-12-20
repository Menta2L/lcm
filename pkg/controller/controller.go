package controller

import (
	"context"
	"crypto/x509"
	"sync"
	"time"

	"github.com/menta2l/lcm/pkg/scheduler"
	log "github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
)

const (
	ControllerName = "issuers"
)

type Controller struct {
	// ctx is the root golang context for the controller
	ctx context.Context

	// the function that should be called when an item is popped
	// off the workqueue
	syncHandler func(ctx context.Context, key string) error

	// queue is a reference to the queue used to enqueue resources
	// to be processed
	queue              workqueue.RateLimitingInterface
	scheduledWorkQueue scheduler.ScheduledWorkQueue
}

func (c *Controller) Register(ctx *context.Context) {
	c.ctx = *ctx
	c.queue = workqueue.NewNamedRateLimitingQueue(DefaultItemBasedRateLimiter(), ControllerName)
	c.scheduledWorkQueue = scheduler.NewScheduledWorkQueue(c.queue.Add)
}
func DefaultItemBasedRateLimiter() workqueue.RateLimiter {
	return workqueue.NewItemExponentialFailureRateLimiter(time.Second*5, time.Minute*5)
}

// Run starts the controller loop
func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	//log := logf.FromContext(ctx)

	log.Info("starting control loop")
	// wait for all the informer caches we depend on are synced
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		// TODO (@munnerz): make time.Second duration configurable
		go wait.Until(func() {
			defer wg.Done()
			c.worker(ctx)
		}, time.Second, stopCh)
	}
	<-stopCh
	log.Info("shutting down queue as workqueue signaled shutdown")
	c.queue.ShutDown()
	log.Info("waiting for workers to exit...")
	wg.Wait()
	log.Info("workers exited")
	return nil
}
func (b *Controller) worker(ctx context.Context) {
	//	log := logf.FromContext(b.ctx)

	log.Info("starting worker")
	for {
		obj, shutdown := b.queue.Get()
		if shutdown {
			break
		}

		var key string
		// use an inlined function so we can use defer
		func() {
			defer b.queue.Done(obj)
			var ok bool
			if key, ok = obj.(string); !ok {
				return
			}
			//log := log.WithValues("key", key)
			log.Info("syncing item")
			if err := b.syncHandler(ctx, key); err != nil {
				log.Error(err, "re-queuing item  due to error processing")
				b.queue.AddRateLimited(obj)
				return
			}
			log.Info("finished processing work item")
			b.queue.Forget(obj)
		}()
	}
	log.Info("exiting worker loop")
}

func (c *Controller) ProcessItem(ctx context.Context, key string) error {
	return c.Sync(ctx)
}
func (c *Controller) Sync(ctx context.Context) (err error) {
	return nil
}

type calculateDurationUntilRenewFn func(context.Context, *x509.Certificate) time.Duration

func scheduleRenewal(ctx context.Context, calc calculateDurationUntilRenewFn, queueFn func(interface{}, time.Duration)) {
}
