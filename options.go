package gorabbit

import amqp "github.com/rabbitmq/amqp091-go"

//////////////////////////////////////////////////////////////////////////////
//								publish
//////////////////////////////////////////////////////////////////////////////

func WithContentType(contentType string) PublishOption {
	return func(publishing *amqp.Publishing) *amqp.Publishing {
		publishing.ContentType = contentType
		return publishing
	}
}

// delay is milliseconds type
func WithDelay(delay int64) PublishOption {
	return func(publishing *amqp.Publishing) *amqp.Publishing {
		if publishing.Headers == nil {
			publishing.Headers = make(amqp.Table)
		}
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

//////////////////////////////////////////////////////////////////////////////
//							exchange and queue
//////////////////////////////////////////////////////////////////////////////

func ExchangeDeclareOption(
	name string,
	kind string,
	durable bool,
	autoDelete bool,
	internal bool,
	noWait bool,
	args amqp.Table,
) Opt {
	return func() interface{} {
		return exchangeDeclareOpt{
			name:       name,
			kind:       kind,
			durable:    durable,
			autoDelete: autoDelete,
			internal:   internal,
			noWait:     noWait,
			args:       args,
		}
	}
}

func QueueDeclareOption(
	name string,
	durable bool,
	autoDelete bool,
	exclusive bool,
	noWait bool,
	args amqp.Table,
) Opt {
	return func() interface{} {
		return queueDeclareOpt{
			name:       name,
			durable:    durable,
			autoDelete: autoDelete,
			exclusive:  exclusive,
			noWait:     noWait,
			args:       args,
		}
	}
}
