package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/upstash/terraform-provider-upstash/upstash"
)

var redisDelete = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a database by id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := upstash.NewUpstashClient(viper.GetString("email"), viper.GetString("key"))

		err := client.DeleteDatabase(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Successfully deleted database %s\n", args[0])
		return nil

	},
}

func init() {

	redisCmd.AddCommand(redisDelete)
}
