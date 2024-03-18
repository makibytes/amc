package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/rc"
	"github.com/makibytes/amc/receive"
	"github.com/spf13/cobra"
)

var getArgs conn.ReceiveArguments
var getCmd = &cobra.Command{
	Use:   "get <queue>",
	Short: "Fetch a message from a queue",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		connArgs = getConnArgs(rootCmd)

		number, _ := cmd.Flags().GetInt("number")
		timeout, _ := cmd.Flags().GetFloat32("timeout")
		wait, _ := cmd.Flags().GetBool("wait")
		if wait {
			timeout = 0
		}

		multicast, _ := cmd.Flags().GetBool("multicast")
		durable, _ := cmd.Flags().GetBool("durable")
		getArgs := conn.ReceiveArguments{
			Acknowledge: true,
			Durable:     durable,
			Multicast:   multicast,
			Number:      number,
			Queue:       args[0],
			Timeout:     timeout,
			Wait:        wait,
		}

		connection, session, err := conn.Connect(connArgs)
		if err != nil {
			return err
		}

		message, err := receive.ReceiveMessage(session, getArgs)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				// no message when timeout occurs? -> no error
				return nil
			} else {
				return err
			}
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
	getCmd.Flags().BoolP("durable", "d", false, "create durable queue if it doesn't exist")
	getCmd.Flags().BoolP("multicast", "m", false, "multicast: subscribe to address, default is anycast: get from queue")
	getCmd.Flags().IntP("number", "n", 1, "number of messages to fetch, 0 = all")
	getCmd.Flags().BoolP("wait", "w", false, "wait (endless) for a message to arrive")
	getCmd.Flags().Float32P("timeout", "t", 0.1, "seconds to wait")
}

func handleMessage(message *amqp.Message, args conn.ReceiveArguments) error {
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
