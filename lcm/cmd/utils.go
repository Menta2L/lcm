package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PrintErrorAndExit logs a fatal error and exits the process with return code 1
func PrintErrorAndExit(err error) {
	log.Fatal(err)
	os.Exit(1)
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
func tableOutput(list []string, c *columnize.Config) string {
	if len(list) == 0 {
		return ""
	}

	delim := "|"
	if c != nil && c.Delim != "" {
		delim = c.Delim
	}

	underline := ""
	headers := strings.Split(list[0], delim)
	for i, h := range headers {
		h = strings.TrimSpace(h)
		u := strings.Repeat("-", len(h))

		underline = underline + u
		if i != len(headers)-1 {
			underline = underline + delim
		}
	}

	list = append(list, "")
	copy(list[2:], list[1:])
	list[1] = underline

	return columnOutput(list, c)
}
func columnOutput(list []string, c *columnize.Config) string {
	if len(list) == 0 {
		return ""
	}

	if c == nil {
		c = &columnize.Config{}
	}
	if c.Glue == "" {
		c.Glue = "    "
	}
	if c.Empty == "" {
		c.Empty = "n/a"
	}

	return columnize.Format(list, c)
}
