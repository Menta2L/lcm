package cmd

import (
	"fmt"
	"github.com/menta2l/lcm/pkg/build"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lcm",
	Long:  `All software has versions. This is lcm's`,
	Run: func(cmd *cobra.Command, args []string) {
		info := build.GetInfo()
		tw := tabwriter.NewWriter(os.Stdout, 2, 1, 2, ' ', 0)
		fmt.Fprintf(tw, "Build Tag:\t%s\n", info.Tag)
		fmt.Fprintf(tw, "Go Version:\t%s\n", info.GoVersion)
		fmt.Fprintf(tw, "Build SHA-1:\t%s\n", info.Revision)
		tw.Flush()
	},
}
