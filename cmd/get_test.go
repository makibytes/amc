package cmd

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Parse without flags", func(t *testing.T) {
		os.Args = []string{"amc", "get", "queue1"}
		err := Execute()
		if err != nil {
			t.Errorf("Error executing command: %v", err)
		}

		if action := getCmd.CalledAs(); action != "get" {
			t.Errorf("Action 'put' has not been called")
		}
		if connArgs.Server != "amqp://localhost:5672" {
			t.Errorf("URI not set to default")
		}
		if getArgs.Queue != "queue1" {
			t.Errorf("Queue name not set correctly: %s", getArgs.Queue)
		}
		// if message := rootCmd.Args(); message != "Hello World" {
		// 	t.Errorf("Message not set correctly: %s", message)
		// }
		if verbose, _ := rootCmd.Flags().GetBool("verbose"); verbose != false {
			t.Errorf("Verbose flag not set correctly: %v", verbose)
		}
	})

}
