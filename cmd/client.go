package cmd

import (
	"time"

	"github.com/menta2l/lcm/pkg/client"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(clientCmd)

	// Flags
	clientCmd.PersistentFlags().StringP("host", "H", "localhost", "The host (server) to connect to")
	clientCmd.PersistentFlags().IntP("port", "p", 6300, "The port the server will listen on")
	clientCmd.PersistentFlags().DurationP("expiry", "e", 30*time.Minute, "How often to check for expired certificates")
	clientCmd.PersistentFlags().DurationP("renewalThreshold", "r", 24*time.Hour, "How early before expiry shall certificates be renewed")

	// Load clientCmd.PersistentFlags() values from config
	bindPrefixedFlags(clientCmd, "client", "host", "port", "storage", "domains", "expiry", "renewalThreshold")
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start a client instance",
	Long:  `Start a client instance, connecting to a running server instance`,
	Run: func(cmd *cobra.Command, args []string) {

		c := config.ClientConfig{
			Hostname:         viper.GetString("client.host"),
			Port:             viper.GetInt("client.port"),
			ExpiryCheckAt:    viper.GetDuration("client.expiry"),
			RenewalThreshold: viper.GetDuration("client.renewalThreshold"),
		}

		client.StartClient(&c, UserAgent())
	},
}
