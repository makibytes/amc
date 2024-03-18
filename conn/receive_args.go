package conn

type ReceiveArguments struct {
	Acknowledge    bool
	Durable        bool
	Multicast      bool
	Number         int
	Queue          string
	Timeout        float32
	Wait           bool
	WithHeader     bool
	WithProperties bool
}
