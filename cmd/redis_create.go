package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/upstash/terraform-provider-upstash/upstash"
)

var redisCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		databaseName := args[0]

		client := upstash.NewUpstashClient(viper.GetString("email"), viper.GetString("key"))

		region, err := cmd.Flags().GetString("region")
		if err != nil {
			return err
		}
		tls, err := cmd.Flags().GetBool("tls")
		if err != nil {
			return err
		}
		consistent, err := cmd.Flags().GetBool("consistent")
		if err != nil {
			return err
		}
		multiZone, err := cmd.Flags().GetBool("multi-zone")
		if err != nil {
			return err
		}

		database, err := client.CreateDatabase(upstash.CreateDatabaseRequest{
			Region:       region,
			DatabaseName: databaseName,
			Tls:          tls,
			Consistent:   consistent,
			MultiZone:    multiZone,
		})
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

	redisCreateCmd.Flags().StringP("region", "r", "", "The region where your databse will be deployed, available regions are: 'eu-west-1', 'us-east-1', 'us-west-1', 'ap-northeast-1' , 'eu-central1'")
	err := redisCreateCmd.MarkFlagRequired("region")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)

	}
	redisCreateCmd.Flags().Bool("tls", false, "This option encrypts the data traffic.")
	redisCreateCmd.Flags().BoolP("consistent", "c", false, "When enabled, writes are guaranteed to be persisted to disk")
	redisCreateCmd.Flags().BoolP("multi-zone", "m", false, "Multi zone replication provides you high availablilty and better scalability")
	redisCmd.AddCommand(redisCreateCmd)

}
