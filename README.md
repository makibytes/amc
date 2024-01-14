# AMC - AMQP 1.0 Messaging Client

[![Build Status](https://travis-ci.org/makibytes/amc.svg?branch=master)](https://travis-ci.org/makibytes/amc)
[![Go Report Card](https://goreportcard.com/badge/github.com/makibytes/amc)](https://goreportcard.com/report/github.com/makibytes/amc)
[![GoDoc](https://godoc.org/github.com/makibytes/amc?status.svg)](https://godoc.org/github.com/makibytes/amc)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](

This project provides a command-line interface (CLI) for sending and receiving messages using the AMQP 1.0 protocol.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.18 or later (see [https://github.com/Azure/go-amqp](Azure/go-amqp))
- An AMQP 1.0 compatible Message Broker, e.g. Azure Message Service or Apache ActiveMQ Artemis

## Usage

You can send a message with the following command:

```sh
amc put <queue-name> <message>
```

The message can also be provided via stdin:

```sh
amc put <queue-name> < message.dat
```

You can receive a message with the following command:

```sh
amc get <queue-name>
```

This will print the payload (data) to stdout and remove the message from the
queue. Use `peek` instead of `get` to keep it in the queue.

The following parameters and environment variables can be used for all commands:

```sh
  -s, --server string     server URL of the AMQP broker  [$AMC_SERVER]
  -p, --password string   password for SASL login        [$AMC_USER]
  -u, --user string       username for SASL login        [$AMC_PASSWORD]
```

## Testing

The tests are based on the [https://https://github.com/bats-core/bats-core](bats testing framework)
(included) and depend on a local Artemis broker with its default settings.

If you have Docker you can spin up an Artemis container like so:

```sh
docker run --name artemis -d \
    -p 8161:8161 -p 5672:5672 \
    apache/activemq-artemis:2.31.2-alpine
```

Port 5672 is the port of the AMQP 1.0 protocol. Port 8161 provides access to the Artemis web console,
where you can check the queues and messages manually. Default credentials are artemis/artemis.

Then you can start the tests:

```sh
./run-tests.sh
```

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
