package send

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

	var durability amqp.Durability
	if args.Durable {
		durability = amqp.DurabilityUnsettledState
	} else {
		durability = amqp.DurabilityNone
	}

	senderOptions := &amqp.SenderOptions{
		Durability: durability,
		//		DynamicAddress:   true,
		SourceAddress:    args.Address,
		TargetDurability: durability,
		Name:             "amc",
	}

	log.Verbose("📤 generating sender...")
	sender, err := session.NewSender(ctx, args.Address, senderOptions)
	if err != nil {
		return err
	}
	defer sender.Close(ctx)

	log.Verbose("💌 sending message...")
	err = sender.Send(ctx, message, nil)

	return err
}
