package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	issuer.AddCommand(addIssuer)
	issuer.AddCommand(listIssuer)
	clientCmd.AddCommand(issuer)

	// Flags
	addIssuer.PersistentFlags().StringP("key", "K", "", "Cloudflare api key")
	addIssuer.PersistentFlags().StringP("sa", "S", "", "Google cloud service account file")
	addIssuer.PersistentFlags().StringP("name", "N", "", "Issuer uniqe name")

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

	},
}

var listIssuer = &cobra.Command{
	Use:   "list",
	Short: "list issueers",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
