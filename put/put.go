package put

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/log"
)

type PutArguments struct {
	Address    string
	Message    []byte
	Durable    bool
	Priority   uint8
	Properties map[string]any
}

// SendMessage sends a message to the specified queue.
func SendMessage(ctx context.Context, session *amqp.Session, args PutArguments) error {
	message := amqp.NewMessage(args.Message)

	//message.Header.Durable = args.Durable
	//message.Header.Priority = args.Priority

	if args.Properties != nil {
		message.ApplicationProperties = args.Properties
	}

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
