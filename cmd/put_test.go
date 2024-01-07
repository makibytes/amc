package cmd

import (
	"os"
	"testing"
)

func TestPut(t *testing.T) {
	t.Run("Parse with flags", func(t *testing.T) {
		os.Args = []string{"amc", "-H", "amqp://custom_uri:1234", "put", "queue1", "-v", "Hello World"}
		err := Execute()
		if err != nil {
			t.Errorf("Error executing command: %v", err)
		}

		if action := putCmd.CalledAs(); action != "put" {
			t.Errorf("Action 'put' has not been called")
		}
		if connArgs.Server != "amqp://custom_uri:1234" {
			t.Errorf("AMQP URI not set correctly: %s", connArgs.Server)
		}
		if putArgs.Address != "queue1" {
			t.Errorf("Address name not set correctly: %s", putArgs.Address)
		}
		if string(putArgs.Message) != "Hello World" {
			t.Errorf("Message not set correctly: %s", putArgs.Message)
		}
		if verbose, _ := rootCmd.Flags().GetBool("verbose"); verbose != true {
			t.Errorf("Verbose flag not set correctly: %v", verbose)
		}
	})

}
