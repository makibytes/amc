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

The following parameters and environment variables can be used:

```sh
  -s, --server string     server URL of the AMQP broker  [$AMC_SERVER]
  -p, --password string   password for SASL login        [$AMC_USER]
  -u, --user string       username for SASL login        [$AMC_PASSWORD]
```

## Running the tests

To run the tests, use the `go test` command:

```sh
go test ./...
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
