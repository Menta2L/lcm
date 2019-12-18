/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/menta2l/lcm"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var certCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		if viper.GetString("cert.name") == "" || viper.GetString("cert.domains") == "" {
			return
		}
		domains := strings.Split(viper.GetString("cert.domains"), ",")
		req := &lcm.CertificateRequest{
			Name:      viper.GetString("cert.name"),
			IssuerRef: viper.GetString("cert.issuer"),
			Domain:    domains,
		}
		client := lcm.GetLcmClient(clientConfig)
		ctx := context.Background()
		_, err := client.RequestCertificate(ctx, req)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	certCmd.AddCommand(certCreateCmd)

	certCreateCmd.PersistentFlags().StringP("name", "N", "", "cert uniqe name")
	certCreateCmd.PersistentFlags().StringP("domains", "D", "", "Comma separated list of domains")
	certCreateCmd.PersistentFlags().StringP("issuer", "I", "", "Issuer name")

	bindPrefixedFlags(certCreateCmd, "cert", "name", "domains", "issuer")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
