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

	"github.com/menta2l/lcm"
	"github.com/menta2l/lcm/pkg/config"
	"github.com/spf13/cobra"
)

var clientConfig config.ClientConfig

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start local certificate manager client",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		client := lcm.GetLcmClient(clientConfig)
		ctx := context.Background()
		req := &lcm.ListSolversRequest{}
		_, err := client.Listsolvers(ctx, req)
		fmt.Println(err)
	},
}

func init() {

	RootCmd.AddCommand(clientCmd)
	clientCmd.PersistentFlags().StringVar(&clientConfig.Host, "H", "127.0.0.1", "service host")
	clientCmd.PersistentFlags().IntVar(&clientConfig.Port, "P", 8000, "service port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
