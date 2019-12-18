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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var solverCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		client := lcm.GetLcmClient(clientConfig)
		ctx := context.Background()
		req := &lcm.Solver{
			Name:  viper.GetString("solver.name"),
			Email: viper.GetString("solver.email"),
		}
		if viper.GetString("solver.key") != "" {
			solver := lcm.Solver_CloudFlareSolver{CloudFlareSolver: &lcm.CloudFlareSolver{Apikey: viper.GetString("solver.key")}}
			req.DnsSolvers = &solver
		}

		_, err := client.CreateSolver(ctx, req)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	solverCmd.AddCommand(solverCreateCmd)
	solverCreateCmd.PersistentFlags().StringP("key", "K", "", "Cloudflare api key")
	solverCreateCmd.PersistentFlags().StringP("sa", "S", "", "Google cloud service account file")
	solverCreateCmd.PersistentFlags().StringP("name", "N", "", "Solver uniqe name")
	solverCreateCmd.PersistentFlags().StringP("email", "E", "", "Eemail address")
	solverCreateCmd.PersistentFlags().StringP("project", "P", "", "Google cloud project id")
	bindPrefixedFlags(solverCreateCmd, "solver", "name", "sa", "key", "email", "project")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
