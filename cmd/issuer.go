package cmd

import (
	"context"

	"github.com/go-acme/lego/v3/log"
	"github.com/menta2l/lcm/pkg/api"
	"github.com/menta2l/lcm/pkg/client"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	issuer.AddCommand(addIssuer)
	issuer.AddCommand(listIssuer)
	clientCmd.AddCommand(issuer)

	// Flags
	addIssuer.PersistentFlags().StringP("key", "K", "", "Cloudflare api key")
	addIssuer.PersistentFlags().StringP("sa", "S", "", "Google cloud service account file")
	addIssuer.PersistentFlags().StringP("name", "N", "", "Issuer uniqe name")
	addIssuer.PersistentFlags().StringP("email", "E", "", "Issuer email")
	addIssuer.PersistentFlags().StringP("project", "P", "", "Google cloud project id")
	bindPrefixedFlags(addIssuer, "issuer", "name", "sa", "key", "email", "project")
}

var issuer = &cobra.Command{
	Use:   "issuer",
	Short: "issuer",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var addIssuer = &cobra.Command{
	Use:   "add",
	Short: "add issueer",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("issuer.email") == "" {
			log.Fatal("email flag is required")
		}
		cfg := config.ClientConfig{
			Hostname:         viper.GetString("client.host"),
			Port:             viper.GetInt("client.port"),
			ExpiryCheckAt:    viper.GetDuration("client.expiry"),
			RenewalThreshold: viper.GetDuration("client.renewalThreshold"),
		}

		c := client.StartClient(&cfg, UserAgent())

		if viper.GetString("issuer.sa") == "" && viper.GetString("issuer.key") == "" {
			log.Infof("Going to create self signed issuer")
			_, err := c.CreateIssuer(context.Background(), &api.IssuerRequest{DnsSolvers: &api.IssuerRequest_SelfSignedIssuer{SelfSignedIssuer: &api.SelfSignedIssuerRequest{Name: "test-issuer"}}})
			if err != nil {
				log.Fatalf("error :%s", err)
			}

		}
		if viper.GetString("issuer.sa") != "" && viper.GetString("issuer.key") != "" {
			log.Fatal("sa and key cant be used togather")
		}

		if viper.GetString("issuer.sa") != "" {
			if viper.GetString("issuer.project") == "" {
				log.Fatal("Project cant be empty when you are using google cloud dns solver")

			}
			acmeIssuer := &api.LetsEncryptIssuerRequest{}
			acmeIssuer.Name = "test-name"
			acmeIssuer.Email = viper.GetString("issuer.email")
			acmeIssuer.Server = "xxxx"
			acmeIssuer.Solver = &api.Solver{DnsSolvers: &api.Solver_GoogleCloudSolver{GoogleCloudSolver: &api.GoogleCloudSolver{Sa: []byte(viper.GetString("issuer.sa")), Email: viper.GetString("issuer.email")}}}
			c.CreateIssuer(context.Background(), &api.IssuerRequest{DnsSolvers: &api.IssuerRequest_AcmeIssuer{AcmeIssuer: acmeIssuer}})
			log.Infof("Going to use lets encrypt with google cloud dns solver")

		}
		if viper.GetString("issuer.key") != "" {
			log.Infof("Going to use lets encrypt with cloud flare solver")

		}

	},
}

var listIssuer = &cobra.Command{
	Use:   "list",
	Short: "list issueers",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
