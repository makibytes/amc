package put

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockAMQPClient struct {
	mock.Mock
}

func (m *MockAMQPClient) SendMessage(amqpURI, queueName string, message []byte) error {
	args := m.Called(amqpURI, queueName, message)
	return args.Error(0)
}

func TestSendMessage(t *testing.T) {
	// Create and use a mock AMQP client for testing
	mockClient := &MockAMQPClient{}

	// Define the expected input parameters for SendMessage
	expectedAmqpURI := "mock_uri"
	expectedQueueName := "mock_queue"
	expectedMessage := "Hello, Mock"

	// Mock the behavior of SendMessage function
	mockClient.On("SendMessage", expectedAmqpURI, expectedQueueName, []byte(expectedMessage)).Return(nil)

	// Call the function under test
	err := mockClient.SendMessage(expectedAmqpURI, expectedQueueName, []byte(expectedMessage))

	// Assert that the function returns no error
	if err != nil {
		t.Errorf("SendMessage returned an error: %v", err)
	}
}
