package conn

type ReceiveArguments struct {
	Acknowledge    bool
	Queue          string
	Timeout        int
	Wait           bool
	WithHeader     bool
	WithProperties bool
}
