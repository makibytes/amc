package cmd

import (
	"os"

	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amc",
	Short: "AMQP 1.0 Messaging Client",
}
var connArgs conn.ConnArguments

// main command for parsing arguments
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)

	var defaultAmqpHost = os.Getenv("AMC_HOST")
	if defaultAmqpHost == "" {
		defaultAmqpHost = "amqp://localhost:5672"
	}

	var defaultSaslUser = os.Getenv("AMC_USER")
	var defaultSaslPassword = os.Getenv("AMC_PASSWORD")

	rootCmd.PersistentFlags().StringP("host", "H", defaultAmqpHost, "URL of the AMQP broker")
	rootCmd.PersistentFlags().StringP("user", "u", defaultSaslUser, "username for SASL login")
	rootCmd.PersistentFlags().StringP("password", "p", defaultSaslPassword, "password for SASL login")

	rootCmd.PersistentFlags().BoolVarP(&log.IsVerbose, "verbose", "v", false, "print verbose output")

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(putCmd)
}

func getConnArgs(cmd *cobra.Command) conn.ConnArguments {
	host, _ := rootCmd.Flags().GetString("host")
	user, _ := rootCmd.Flags().GetString("user")
	password, _ := rootCmd.Flags().GetString("password")

	connArgs = conn.ConnArguments{
		Host:     host,
		User:     user,
		Password: password,
	}

	return connArgs
}
