package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lileio/lile/v2"
	"github.com/lileio/pubsub/v2"
	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/config"
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

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			lile.Run()
		}()
		go func() {
			pubsub.Subscribe(&subscribers.LcmServiceSubscriber{})
		}()

		<-c
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
