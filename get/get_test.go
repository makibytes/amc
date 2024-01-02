package get

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockAMQPClient struct {
	mock.Mock
}

func (m *MockAMQPClient) ReceiveMessage(amqpURI, queueName string) ([]byte, error) {
	args := m.Called(amqpURI, queueName)
	return args.Get(0).([]byte), args.Error(1)
}

func TestReceiveMessage(t *testing.T) {
	// Create and use a mock AMQP client for testing
	mockClient := &MockAMQPClient{}

	// Define the expected input parameters for ReceiveMessage
	expectedAmqpURI := "mock_uri"
	expectedQueueName := "mock_queue"

	// Mock the behavior of ReceiveMessage function
	mockClient.On("ReceiveMessage", expectedAmqpURI, expectedQueueName).Return([]byte("Mock Message"), nil)

	// Call the function under test
	message, err := mockClient.ReceiveMessage(expectedAmqpURI, expectedQueueName)
	if err != nil {
		t.Errorf("Received error: %v", err)
	}

	// Add your assertions here
	// Assert that the function returns the expected message content
	if err != nil {
		t.Errorf("ReceiveMessage returned an error: %v", err)
	}
	if string(message) != "Mock Message" {
		t.Errorf("Received message content does not match")
	}
}
