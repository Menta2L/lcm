package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/apex/log"
	"github.com/lileio/lile/v2"
	"github.com/lileio/pubsub/v2"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/menta2l/lcm/pkg/controller"
	"github.com/menta2l/lcm/pkg/util"
	"github.com/menta2l/lcm/pkg/vault"
	"github.com/menta2l/lcm/server"
	"github.com/menta2l/lcm/subscribers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC service",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		wg.Add(1)
		cfg := config.VaultConfig{
			Address: viper.GetString("vault.address"),
			Token:   viper.GetString("vault.token"),
		}
		vc, err := vault.NewClient(cfg)
		if err != nil {
			fmt.Println(err)
			return
		}
		s := &server.LcmServer{Backend: vc}
		lile.Name("lcm")
		lile.Server(func(g *grpc.Server) {
			lcm.RegisterLcmServer(g, s)
		})

		//c := make(chan os.Signal, 1)
		stopCh := SetupSignalHandler()
		rootCtx := util.ContextWithStopCh(context.Background(), stopCh)
		//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		iface := &controller.Controller{}
		iface.Register(&rootCtx)
		//rootCtx := util.ContextWithStopCh(context.Background(), c)
		go func(fn *controller.Controller) {
			defer wg.Done()
			log.Info("starting controller")

			workers := 5
			err := fn.Run(workers, stopCh)
			if err != nil {
				log.Error("error starting controller")
				os.Exit(1)
			}
		}(iface)

		go func() {
			lile.Run()
		}()
		go func() {
			pubsub.Subscribe(&subscribers.LcmServiceSubscriber{})
		}()

		<-stopCh
		lile.Shutdown()
		pubsub.Shutdown()
	},
}

func init() {
	upCmd.PersistentFlags().StringP("address", "", "http://127.0.0.1:8200", "Vault address")
	upCmd.PersistentFlags().StringP("token", "", "s.ADwFgkioiRwInzs3HFbBW7TK", "Vault token")
	bindPrefixedFlags(upCmd, "vault", "address", "token")

	RootCmd.AddCommand(upCmd)
}

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
