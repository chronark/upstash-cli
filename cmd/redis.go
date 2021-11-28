package cmd

import (
	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "This command does nothing by itself, please use a subcommand",
}

func init() {
	rootCmd.AddCommand(redisCmd)
}
