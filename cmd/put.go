package cmd

import (
	"context"
	"io"
	"os"

	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/put"
	"github.com/spf13/cobra"
)

func init() {
	putCmd.Flags().BoolP("multicast", "m", false, "send to a multicast address")
}

var putArgs put.PutArguments
var putCmd = &cobra.Command{
	Use:   "put <address> <message>",
	Short: "Send a message to an address",
	Args:  cobra.MinimumNArgs(1), // message can be read from stdin
	RunE: func(cmd *cobra.Command, args []string) error {
		connArgs = getConnArgs(rootCmd)

		var data []byte
		if len(args) > 1 {
			data = []byte(args[1])
		} else {
			var err error
			data, err = dataFromStdin()
			if err != nil {
				return err
			}
		}

		putArgs = put.PutArguments{
			Address:    args[0],
			Message:    data,
			Durable:    true,
			Priority:   0,
			Properties: nil,
		}

		connection, session, err := conn.Connect(connArgs)
		if err != nil {
			return err
		}

		err = put.SendMessage(context.Background(), session, putArgs)

		session.Close(context.Background())
		connection.Close()

		return err
	},
}

func dataFromStdin() ([]byte, error) {
	f := os.Stdin
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}
