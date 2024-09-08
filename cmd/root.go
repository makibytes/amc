package cmd

import (
	"os"

	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/log"
	"github.com/spf13/cobra"
)

// need a field in order to override it at link time (github action)
var version = "dev"

var rootCmd = &cobra.Command{
	Use:          "amc",
	Short:        "Artemis Messaging Client",
	SilenceUsage: true, // for errors other than wrong command line
	Version:      version,
}
var connArgs conn.ConnArguments

// main command for parsing arguments
func Execute() error {
	return rootCmd.Execute()
}

/*
 * convention:
 * - use lowercase single character flags for connection settings
 * - common connection flags should also be accessible by environment variables
 */
func init() {
	var defaultAmqpServer = os.Getenv("AMC_SERVER")
	if defaultAmqpServer == "" {
		defaultAmqpServer = "amqp://localhost:5672"
	}

	var defaultSaslUser = os.Getenv("AMC_USER")
	var defaultSaslPassword = os.Getenv("AMC_PASSWORD")

	rootCmd.PersistentFlags().StringP("server", "s", defaultAmqpServer, "server URL")
	rootCmd.PersistentFlags().StringP("user", "u", defaultSaslUser, "user for SASL PLAIN login, otherwise ANONYMOUS login)")
	rootCmd.PersistentFlags().StringP("password", "p", defaultSaslPassword, "password for SASL PLAIN login")

	rootCmd.PersistentFlags().BoolVarP(&log.IsVerbose, "verbose", "v", false, "print verbose output")
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)

	rootCmd.AddCommand(peekCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(putCmd)
}

func getConnArgs() conn.ConnArguments {
	server, _ := rootCmd.Flags().GetString("server")
	user, _ := rootCmd.Flags().GetString("user")
	password, _ := rootCmd.Flags().GetString("password")

	connArgs = conn.ConnArguments{
		Server:   server,
		User:     user,
		Password: password,
	}

	return connArgs
}
