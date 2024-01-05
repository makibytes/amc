package conn

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/makibytes/amc/log"
)

type ConnArguments struct {
	Host     string
	User     string
	Password string
}

func Connect(args ConnArguments) (*amqp.Conn, *amqp.Session, error) {
	ctx := context.WithoutCancel(context.Background())

	var connOptions *amqp.ConnOptions
	if args.User == "" {
		connOptions = &amqp.ConnOptions{
			SASLType: amqp.SASLTypeAnonymous(),
		}
	} else {
		connOptions = &amqp.ConnOptions{
			SASLType: amqp.SASLTypePlain(args.User, args.Password),
		}
	}

	log.Verbose("📡 connecting to %s...\n", args.Host)
	connection, err := amqp.Dial(ctx, args.Host, connOptions)
	if err != nil {
		return nil, nil, err
	}

	session, err := connection.NewSession(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return connection, session, nil
}
