package get

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/conn"
	"github.com/makibytes/amc/log"
)

func ReceiveMessage(ctx context.Context, session *amqp.Session, args conn.ReceiveArguments) (*amqp.Message, error) {
	log.Verbose("📥 generating receiver...")
	receiver, err := session.NewReceiver(ctx, args.Queue, nil)
	if err != nil {
		return nil, err
	}
	defer receiver.Close(ctx)

	log.Verbose("📩 calling receive()...")
	message, err := receiver.Receive(ctx, nil)
	if err != nil {
		return nil, err
	}

	// get: yes, peek: no
	if args.Acknowledge {
		receiver.AcceptMessage(ctx, message)
	}

	return message, nil
}
