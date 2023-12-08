package pkg

import amqp "github.com/rabbitmq/amqp091-go"

func WithContentType(contentType string) PublishOption {
	return func(publishing *amqp.Publishing) *amqp.Publishing {
		publishing.ContentType = contentType
		return publishing
	}
}

// delay is milliseconds type
func WithDelay(delay int64) PublishOption {
	return func(publishing *amqp.Publishing) *amqp.Publishing {
		publishing.Headers["x-delay"] = delay
		return publishing
	}
}

func WithPriority(priority uint8) PublishOption {
	return func(publishing *amqp.Publishing) *amqp.Publishing {
		publishing.Priority = priority
		return publishing
	}
}
