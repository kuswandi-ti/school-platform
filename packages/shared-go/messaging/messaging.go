package messaging

import "context"

type Message struct {
	RoutingKey    string
	Body          []byte
	Headers       map[string]string
	RequestID     string
	CorrelationID string
}

type Publisher interface {
	Publish(ctx context.Context, message Message) error
}

type Handler interface {
	Handle(ctx context.Context, message Message) error
}
