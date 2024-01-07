package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/put"
	"github.com/spf13/cobra"
)

func init() {
	putCmd.Flags().StringP("contenttype", "C", "text/plain", "MIME type of message data")
	putCmd.Flags().StringP("correlationid", "c", "", "correlation id for request/response")
	putCmd.Flags().BoolP("durable", "d", false, "message is durable (stored to disk)")
	putCmd.Flags().StringP("messageid", "i", "", "message id")
	putCmd.Flags().BoolP("multicast", "m", false, "send to a multicast address")
	putCmd.Flags().Uint8("priority", 4, "priority of the message (0-9)")
	putCmd.Flags().StringSliceP("property", "P", []string{}, "message properties in key=value format (can be used multiple times)")
	putCmd.Flags().StringP("replyto", "r", "", "reply to address for request/response")
	putCmd.Flags().String("subject", "", "subject")
	putCmd.Flags().String("to", "", "intended destionation node")
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

		contenttype, _ := cmd.Flags().GetString("contenttype")
		correlationid, _ := cmd.Flags().GetString("correlationid")
		durable, _ := cmd.Flags().GetBool("durable")
		messageid, _ := cmd.Flags().GetString("messageid")
		multicast, _ := cmd.Flags().GetBool("multicast")
		priority, _ := cmd.Flags().GetUint8("priority")
		replyto, _ := cmd.Flags().GetString("replyto")
		subject, _ := cmd.Flags().GetString("subject")
		to, _ := cmd.Flags().GetString("to")

		properties := make(map[string]any) // Initialize the properties map
		propertySlice, _ := cmd.Flags().GetStringSlice("property")
		for _, property := range propertySlice {
			keyValue := strings.SplitN(property, "=", 2)
			if len(keyValue) == 2 {
				properties[keyValue[0]] = keyValue[1]
			} else {
				return fmt.Errorf("invalid property: %s", property)
			}
		}

		putArgs = put.PutArguments{
			Address:       args[0],
			ContentType:   contenttype,
			CorrelationID: correlationid,
			Durable:       durable,
			Message:       data,
			MessageID:     messageid,
			Multicast:     multicast,
			Priority:      priority,
			Properties:    properties, // Assign the properties map
			ReplyTo:       replyto,
			Subject:       subject,
			To:            to,
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
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, errors.New("message missing")
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	return data, nil
}
