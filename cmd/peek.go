package cmd

import (
	"context"
	"errors"

	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/rc"
	"github.com/makibytes/amc/receive"
	"github.com/spf13/cobra"
)

var peekArgs conn.ReceiveArguments
var peekCmd = &cobra.Command{
	Use:   "peek <queue>",
	Short: "Look into a message, but let it stay in the queue",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		connArgs = getConnArgs()

		number, _ := cmd.Flags().GetInt("number")
		timeout, _ := cmd.Flags().GetFloat32("timeout")
		wait, _ := cmd.Flags().GetBool("wait")
		if wait {
			timeout = 0
		}

		durable, _ := cmd.Flags().GetBool("durable")
		peekArgs = conn.ReceiveArguments{
			Acknowledge: false,
			Durable:     durable,
			Number:      number,
			Queue:       args[0],
			Timeout:     timeout,
			Wait:        wait,
		}

		connection, session, err := conn.Connect(connArgs)
		if err != nil {
			return err
		}

		message, err := receive.ReceiveMessage(session, peekArgs)
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

		// use cmd.get's handleMessage()
		err = handleMessage(message, peekArgs)

		session.Close(context.Background())
		connection.Close()

		return err
	},
}

func init() {
	peekCmd.Flags().BoolP("durable", "d", true, "create durable queue if it doesn't exist")
	peekCmd.Flags().IntP("number", "n", 1, "number of messages to fetch, 0 = all")
	peekCmd.Flags().BoolP("wait", "w", false, "wait (endless) for a message to arrive")
	peekCmd.Flags().Float32P("timeout", "t", 0.1, "seconds to wait")
}
