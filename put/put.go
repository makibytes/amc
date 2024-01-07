package put

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/log"
)

type PutArguments struct {
	Address       string
	ContentType   string
	CorrelationID string
	Durable       bool
	Message       []byte
	MessageID     string
	Multicast     bool
	Priority      uint8
	Properties    map[string]any
	ReplyTo       string
	Subject       string
	To            string
}

// SendMessage sends a message to the specified queue.
func SendMessage(ctx context.Context, session *amqp.Session, args PutArguments) error {
	message := amqp.NewMessage(args.Message)

	message.Header.Durable = args.Durable
	message.Header.Priority = args.Priority

	message.Properties.ContentType = &args.ContentType
	message.Properties.CorrelationID = &args.CorrelationID
	message.Properties.MessageID = &args.MessageID
	message.Properties.ReplyTo = &args.ReplyTo
	message.Properties.Subject = &args.Subject
	message.Properties.To = &args.To

	if len(args.Properties) > 0 {
		message.ApplicationProperties = args.Properties
	}

	// if args.Multicast {
	// }

	log.Verbose("📤 generating sender...")
	sender, err := session.NewSender(ctx, args.Address, nil)
	if err != nil {
		return err
	}
	defer sender.Close(ctx)

	log.Verbose("💌 sending message...")
	err = sender.Send(ctx, message, nil)

	return err
}
