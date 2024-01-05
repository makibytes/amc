package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/get"
	"github.com/makibytes/amc/rc"
	"github.com/spf13/cobra"
)

var getArgs get.GetArguments
var getCmd = &cobra.Command{
	Use:   "get <queue>",
	Short: "Fetch a message from a queue",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		connArgs = getConnArgs(rootCmd)

		wait, _ := cmd.Flags().GetBool("wait")
		timeout, _ := cmd.Flags().GetInt("timeout")

		getArgs = get.GetArguments{
			Queue:   args[0],
			Timeout: timeout,
			Wait:    wait,
		}

		connection, session, err := conn.Connect(connArgs)
		if err != nil {
			return err
		}

		timeout, _ = cmd.Flags().GetInt("timeout")
		wait, _ = cmd.Flags().GetBool("wait")
		if wait {
			timeout = 0 // wait flag overrides timeout
		}

		var ctx context.Context
		var cancel context.CancelFunc
		if timeout == 0 {
			ctx, cancel = context.WithCancel(context.Background())
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		}
		defer cancel()

		message, err := get.ReceiveMessage(ctx, session, getArgs)
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
	getCmd.Flags().BoolP("wait", "w", false, "wait (endless) for a message to arrive")
	getCmd.Flags().Int32P("timeout", "t", 30, "seconds to wait")
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
