package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/upstash/terraform-provider-upstash/upstash"
)

var redisGet = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a specific database by id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := upstash.NewUpstashClient(viper.GetString("email"), viper.GetString("key"))

		database, err := client.GetDatabase(args[0])
		if err != nil {
			return err
		}

		b, err := json.MarshalIndent(database, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil

	},
}

func init() {

	redisCmd.AddCommand(redisGet)
}
