package get

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/log"
)

type GetArguments struct {
	Queue          string
	WithHeader     bool
	WithProperties bool
}

// ReceiveMessage receives a message from the specified queue.
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

	return message, nil
}
