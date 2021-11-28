package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `The Upstash cli requires email and an api key in order to authenticate.
Email is your email address while registering Upstash. 
API KEY can be generated from the Upstash Console.
Plese see our API KEY documentation.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := viper.SafeWriteConfig()
		if err != nil {
			return fmt.Errorf("%w\nPlease remove it manually and run this command again\n", err)
		}
		fmt.Printf("Succesfully stored credentials in %s", viper.ConfigFileUsed())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().String("email", "", "The email you used to sign up with on upstash.com")
	authCmd.Flags().String("key", "", "The api key can be generated from the Upstash Console")

	err := authCmd.MarkFlagRequired("email")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)

	}
	err = authCmd.MarkFlagRequired("key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	err = viper.BindPFlag("email", authCmd.Flags().Lookup("email"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	err = viper.BindPFlag("key", authCmd.Flags().Lookup("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

}
