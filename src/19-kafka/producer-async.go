package main

import (
    "os"
    "github.com/Shopify/sarama"
    "strings"
    "log"
    "sync"
    "os/signal"
    "time"
)

var (
    kafka = "172.17.0.8:9092,172.17.0.7:9092,172.17.0.6:9092"
    index = 1
)
func main() {

    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    producer, err := sarama.NewAsyncProducer(strings.Split(kafka, ","), config)
    if err != nil {
        panic(err)
    }

    // Trap SIGINT to trigger a graceful shutdown.
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    var (
        wg                          sync.WaitGroup
        enqueued, successes, errors int
    )

    wg.Add(1)
    go func() {
        defer wg.Done()
        for range producer.Successes() {
            successes++
        }
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for err := range producer.Errors() {
            log.Println(err)
            errors++
        }
    }()

    //ProducerLoop:
    for {
        index++
        message := &sarama.ProducerMessage{Topic: "user", Value: sarama.StringEnn
coder("testing 123")}
        select {
        case producer.Input() <- message:
            enqueued++
            if index > 5 {
                time.Sleep(time.Duration(2) * time.Second)
                producer.AsyncClose() // Trigger a shutdown of the producer.
                break
            }

        case <-signals:
            producer.AsyncClose() // Trigger a shutdown of the producer.
            break
        }
        time.Sleep(time.Duration(2) * time.Second)
    }

    wg.Wait()

    log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}