package main

import (
	"os"
	"testing"

	"github.com/makibytes/amc/cmd"
	"github.com/stretchr/testify/assert"
)

func TestMain_PutAction(t *testing.T) {
	os.Args = []string{"amc", "put", "queue1", "Hello World", "-uri", "amqp://localhost:5672"}
	rc := cmd.Execute()

	assert.Equal(t, nil, rc)
}

func TestMain_GetAction(t *testing.T) {
	os.Args = []string{"amc", "get", "queue1", "-uri", "amqp://localhost:5672"}
	rc := cmd.Execute()

	assert.Equal(t, nil, rc)
}
