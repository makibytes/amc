package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/log"
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
		connArgs = getConnArgs()

		number, _ := cmd.Flags().GetInt("number")
		timeout, _ := cmd.Flags().GetFloat32("timeout")
		wait, _ := cmd.Flags().GetBool("wait")
		if wait {
			timeout = 0
		}

		multicast, _ := cmd.Flags().GetBool("multicast")
		durable, _ := cmd.Flags().GetBool("durable")

		withHeaderAndProperties := log.IsVerbose
		withApplicationProperties, _ := cmd.Flags().GetBool("quiet")
		withApplicationProperties = !withApplicationProperties || log.IsVerbose

		getArgs := conn.ReceiveArguments{
			Acknowledge:               true,
			Durable:                   durable,
			Multicast:                 multicast,
			Number:                    number,
			Queue:                     args[0],
			Timeout:                   timeout,
			Wait:                      wait,
			WithHeaderAndProperties:   withHeaderAndProperties,
			WithApplicationProperties: withApplicationProperties,
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
	getCmd.Flags().Float32P("timeout", "t", 0.1, "seconds to wait")
	getCmd.Flags().BoolP("quiet", "q", false, "quiet about properties, show data only")
	getCmd.Flags().BoolP("wait", "w", false, "wait (endless) for a message to arrive")
}

func handleMessage(message *amqp.Message, args conn.ReceiveArguments) error {
	if args.WithHeaderAndProperties {
		fmt.Fprintf(os.Stderr, "Header: %+v\n", message.Header)
		fmt.Fprintf(os.Stderr, "MessageProperties: %+v\n", message.Properties)
	}
	if args.WithApplicationProperties && len(message.ApplicationProperties) > 0 {
		propertiesString := ""
		for k, v := range message.ApplicationProperties {
			if propertiesString != "" {
				propertiesString += ","
			}
			propertiesString += fmt.Sprintf("%s=%s", k, v)
		}
		fmt.Fprintf(os.Stderr, "Properties: %s\n", propertiesString)
	}

	// always print message data
	fmt.Print(string(message.GetData()))
	// add newline if stdout just for better readability
	if log.IsStdout {
		fmt.Println()
	}

	return nil
}
