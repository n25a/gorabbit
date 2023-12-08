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

### 🐇 Create a connection

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

### ✍️ Declare a queue and a exchange

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
	
    // Create a consumer
    err = rabbit.StartJobs()
    if err != nil {
        panic(err)
    }
}
```

### 📨 Publish a job

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

