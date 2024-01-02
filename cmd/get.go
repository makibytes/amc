package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/get"
	"github.com/makibytes/amc/rc"
	"github.com/spf13/cobra"
)

var getArgs get.GetArguments
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch a message from a queue",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		connArgs = getConnArgs(rootCmd)
		getArgs = get.GetArguments{
			Queue: args[0],
		}

		connection, session, err := conn.Connect(connArgs)
		if err != nil {
			return err
		}

		message, err := get.ReceiveMessage(context.Background(), session, getArgs)
		if err != nil {
			return err
		}
		if message == nil {
			return errors.New(rc.NoMessage)
		}

		err = handleMessage(message, getArgs)

		session.Close(context.Background())
		connection.Close()

		return err
	},
}

func init() {
	getCmd.Flags().Int32P("number", "n", 1, "number of messages to fetch, 0 = all")
}

func handleMessage(message *amqp.Message, args get.GetArguments) error {
	if args.WithHeader {
		headerString := fmt.Sprintf("%v", message.Header)
		fmt.Printf("Header:\n%s\n", headerString)
	}
	if args.WithProperties {
		propertiesString := fmt.Sprintf("%v", message.Properties)
		fmt.Printf("Header:\n%s\n", propertiesString)
	}
	if args.WithHeader || args.WithProperties {
		fmt.Println("Data:")
	}

	// always print message data
	fmt.Println(string(message.GetData()))

	return nil
}
