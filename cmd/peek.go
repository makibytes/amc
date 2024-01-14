package cmd

import (
	"context"
	"errors"
	"time"

	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/get"
	"github.com/makibytes/amc/rc"
	"github.com/spf13/cobra"
)

var peekArgs conn.ReceiveArguments
var peekCmd = &cobra.Command{
	Use:   "peek <queue>",
	Short: "Look into a message, but let it stay in the queue",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		connArgs = getConnArgs(rootCmd)

		wait, _ := cmd.Flags().GetBool("wait")
		timeout, _ := cmd.Flags().GetInt("timeout")

		peekArgs = conn.ReceiveArguments{
			Acknowledge: false,
			Queue:       args[0],
			Timeout:     timeout,
			Wait:        wait,
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

		message, err := get.ReceiveMessage(ctx, session, peekArgs)
		if err != nil {
			return err
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
	peekCmd.Flags().Int32P("number", "n", 1, "number of messages to fetch, 0 = all")
	peekCmd.Flags().BoolP("wait", "w", false, "wait (endless) for a message to arrive")
	peekCmd.Flags().Int32P("timeout", "t", 30, "seconds to wait")
}
