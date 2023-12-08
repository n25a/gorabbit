![logo](https://raw.githubusercontent.com/n25a/gorabbit/master/docs/logo.jpg)


# 🐇 gorabbit - RabbitMQ client for Golang

[![GoDoc](https://godoc.org/github.com/n25a/gorabbit?status.svg)](https://godoc.org/github.com/n25a/gorabbit)
![License](https://img.shields.io/badge/license-MIT-blue.svg)


## 📖 Introduction

gorabbit is a RabbitMQ client for Golang. It is a wrapper around [rabbitmq/amqp091-go](github.com/rabbitmq/amqp091-go) library. 
It provides a simple interface to interact with RabbitMQ. Also, it provides a simple way to create a consumer and a publisher that we call them `jobs`.
It's good to mentioned that gorabbit handles the reconnection and reconsuming the jobs automatically.


## 📦 Installation

```bash
go get github.com/n25a/gorabbit
```


## ⚙️ How to use

In the following, we will show you how to use gorabbit in your project.

### 🐇 Create a connection

At the first step, you should create a rabbitmq instance. Then, you can create a connection to the rabbitmq server.

```go
import (
    "github.com/n25a/gorabbit"
)

func main() {
    // Create a rabbitmq instance 
    rabbit := gorabbit.NewRabbitMQ(
        dsn,
        dialTimeout,
        dialRetry,
        ctxTimeout,
        logLevel,
    ) 
    
    // Create a connection
    err := rabbit.Connect()
    if err != nil {
        panic(err)
    }
    
    // Close the connection
    defer rabbit.Close()
}
```

Each parameter of the `NewRabbitMQ` function is described below:  
* dsn: RabbitMQ connection string. It should be in the following format: `amqp://user:password@host:port/vhost`
* dialTimeout: The timeout for dialing to the RabbitMQ server. The default value is 5 seconds.
* dialRetry: The number of retries for dialing to the RabbitMQ server. The default value is 0.
* ctxTimeout: The timeout for the context of consumer handler. The default value is 1 second.
* logLevel: The log level for the gorabbit. The default value is `info`. It can be `debug`, `info`, `warn`, `error`, `fatal`, `panic`, `dpanic`.

This function create instance of the `RabbitMQ` struct. This struct has the following methods:
* Connect: Create a connection to the RabbitMQ server.
* Close: Close the connection to the RabbitMQ server.
* Declare: Declare a queue and an exchange.
* StartConsumingJobs: Start the consumers.
* NewJob: Create a new job instance.
* ShutdownJobs: Shutdown the consumers.
* GetConnection: Get the connection to the RabbitMQ server.
* GetChannel: Get the channel to the RabbitMQ server.

For connecting to the RabbitMQ server, you should call the `Connect` method. Also, you should call the `Close` method for closing the connection.


### ✍️ Declare a queue and a exchange

For declaring a queue and an exchange, you should call the `Declare` method. This method has the following parameters:
* exchangeOption: The exchange options. It can be created by the `ExchangeDeclareOption` function.
* queueOption: The queue options. It can be created by the `QueueDeclareOption` function. It's good to mentioned that you can pass multiple queue options to this function for same exchange.

```go
import (
    "github.com/n25a/gorabbit"
)

func main() {
    //...
    
    // Close the connection
    defer rabbit.Close()
    
    
    // Declare an exchange 
    exchangeOption := gorabbit.ExchangeDeclareOption(
        name,
        kind,
        durable,
        autoDelete,
        internal,
        noWait,
        args,
    )
	
    // Declare a queue
    queueOption := gorabbit.QueueDeclareOption(
        name,
        durable,
        autoDelete,
        exclusive,
        noWait,
        args,
    )
    
    // Declare a queue and an exchange for rabbit instance
    err = rabbit.Declare(exchangeOption, queueOption)
    if err != nil {
        panic(err)
    }
}
```

### 📩 Consume a job

For consuming a job, you should create a job instance. Then, you should call the `StartConsumingJobs` method for starting the consumers.
This method, start consuming of all jobs that created by `gorabbit.RabbitMQ` instance.

Creating new `job` instance has the following parameters:
* handler: The handler function for consuming the job. It should be in the following format: `func(ctx context.Context, message []byte) error`.
* jobExchange: The exchange name for consuming the job.
* jobQueue: The queue name for consuming the job.
* autoAck: The auto ack for consuming the job.
* justPublish: It's a flag for just publishing the job. It's useful for start all consumers and pass publishers jobs.

```go
import (
    "github.com/n25a/gorabbit"
)

func main() {
    //...
    
    // Declare a queue and an exchange for rabbit instance
    err = rabbit.Declare(exchangeOption, queueOption)
    if err != nil {
        panic(err)
    }
	
    // create a job instance
	job := rabbit.NewJob(
        handler,
        jobExchange,
        jobQueue,
        autoAck,
        justPublish,
    )
	
    // Create a consumers
    err = rabbit.StartConsumingJobs()
    if err != nil {
        panic(err)
    }
}
```

### 📨 Publish a job

For publishing a job, you should create a job instance. Then, you should call the `Publish` method for publishing the job.
This method, publish the message on declared job to the RabbitMQ server.
The `Publish` method has the following parameters:
* ctx: The context for publishing the job.
* message: The message for publishing the job. It should be in the `[]byte` format.
* options: The options for publishing the job. It can be created by the `PublishOption` function. These optional options are described below:
    * WithContentType: The content type for publishing the job. The default value is `text/json`.
    * WithDelay: The delay for publishing the job. The default value is 0.
    * WithPriority: The priority for publishing the job. The default value is 0.

```go
import (
    "github.com/n25a/gorabbit"
)

func main() {
    //...

    // Create a consumer
    err = rabbit.StartJobs()
    if err != nil {
        panic(err)
    }
	
    // Create a publisher
    msg, _ := json.Marshal(struct{
        Body string `json:"body"`
    }{
        Body: "test",
    })
	
    err = job.Publish(
        context.Background(),
        msg,
        gorabbit.WithContentType("text/json"),
        gorabbit.WithDelay(10),
        gorabbit.WithPriority(0),
    )
    if err != nil {
        panic(err)
    }
}
```