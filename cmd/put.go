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

/*
 * convention:
 * - use uppercase flags for message creation
 * - not all flags need to have a single char flag
 * - common flags should have a single char flag of course
 * (let's see how long we can keep this up)
 */
func init() {
	putCmd.Flags().StringP("contenttype", "T", "text/plain", "MIME type of message data")
	putCmd.Flags().StringP("correlationid", "C", "", "correlation id for request/response")
	putCmd.Flags().BoolP("durable", "D", false, "message is durable (stored to disk)")
	putCmd.Flags().StringP("messageid", "I", "", "message id")
	putCmd.Flags().BoolP("multicast", "M", false, "send to a multicast address, default is anycast")
	putCmd.Flags().Uint8P("priority", "Y", 4, "priority of the message (0-9)")
	putCmd.Flags().StringSliceP("property", "P", []string{}, "message properties in key=value format (can be used multiple times)")
	putCmd.Flags().StringP("replyto", "R", "", "reply to address for request/response")
	putCmd.Flags().String("subject", "", "subject")
	putCmd.Flags().String("to", "", "intended destionation node")
}

var putArgs conn.SendArguments
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

		putArgs = conn.SendArguments{
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
