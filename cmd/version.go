package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lcm",
	Long:  `All software has versions. This is lcm's`,
	Run: func(cmd *cobra.Command, args []string) {
		versionString := fmt.Sprintf("lcm %s (%s)", Version, Commit)
		fmt.Println(versionString)
	},
}
