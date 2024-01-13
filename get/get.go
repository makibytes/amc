package get

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/log"
)

type GetArguments struct {
	Acknowledge    bool
	Queue          string
	Timeout        int
	Wait           bool
	WithHeader     bool
	WithProperties bool
}

func ReceiveMessage(ctx context.Context, session *amqp.Session, args GetArguments) (*amqp.Message, error) {
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
