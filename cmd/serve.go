package cmd

import (
	"errors"
	"time"

	"github.com/go-acme/lego/v3/log"

	"github.com/menta2l/lcm/pkg/config"
	"github.com/menta2l/lcm/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)

	// Flags
	serveCmd.PersistentFlags().IntP("port", "p", 6300, "The port the server will listen on")
	serveCmd.PersistentFlags().DurationP("expiry", "x", 30*time.Minute, "How often to check for expired certificates")
	serveCmd.PersistentFlags().DurationP("renewalThreshold", "r", (24*time.Hour)*15, "How early before expiry shall certificates be renewed")

	// Load serveCmd.PersistentFlags() values from config
	bindPrefixedFlags(serveCmd, "server", "port", "expiry", "renewalThreshold")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a server instance",
	Long:  `Start the server component, doing the interaction with the ACME server and connected clients`,
	Run: func(cmd *cobra.Command, args []string) {

		c := config.ServerConfig{
			Port:             viper.GetInt("server.port"),
			ExpiryCheckAt:    viper.GetDuration("server.expiry"),
			RenewalThreshold: viper.GetDuration("server.renewalThreshold"),
		}

		if c.RenewalThreshold > (24*time.Hour)*60 {
			PrintErrorAndExit(errors.New("renewal threshold can't exceed 60 days"))
		}

		server.StartServer(&c, UserAgent())
		log.Infof("Listening on port %v", c.Port)
	},
}

func bindPrefixedFlag(cmd *cobra.Command, prefix string, key string) {
	err := viper.BindPFlag(prefix+"."+key, cmd.PersistentFlags().Lookup(key))

	if err != nil {
		PrintErrorAndExit(err)
	}
}

func bindPrefixedFlags(cmd *cobra.Command, prefix string, keys ...string) {
	for i := 0; i < len(keys); i++ {
		bindPrefixedFlag(cmd, prefix, keys[i])
	}
}
