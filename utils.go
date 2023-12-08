package gorabbit

import amqp "github.com/rabbitmq/amqp091-go"

type Opt func() interface{}

type PublishOption func(p *amqp.Publishing) *amqp.Publishing

type exchangeDeclareOpt struct {
	name       string
	kind       string
	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
	args       amqp.Table
}

type queueDeclareOpt struct {
	name       string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	args       amqp.Table
}

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
