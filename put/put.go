package put

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/log"
)

func SendMessage(ctx context.Context, session *amqp.Session, args conn.SendArguments) error {
	log.Verbose("🏗️ constructing message...")
	message := amqp.NewMessage(args.Message)
	message.Header = &amqp.MessageHeader{
		Durable:  args.Durable,
		Priority: args.Priority,
	}
	message.Properties = &amqp.MessageProperties{
		ContentType:   &args.ContentType,
		CorrelationID: &args.CorrelationID,
		MessageID:     &args.MessageID,
		ReplyTo:       &args.ReplyTo,
		Subject:       &args.Subject,
		To:            &args.To,
	}

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
