package main 

import (
	"github.com/Shopify/sarama"
	"os"
	"strings"
	"log"
)

var (
    kafka = "172.17.0.8:9092,172.17.0.7:9092,172.17.0.6:9092"
    logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger
	// 连接kafka消息服务器
    producer, err := sarama.NewSyncProducer(strings.Split(kafka, ","), nil)
    if err != nil {
        log.Fatalln(err)
    }
    defer func() {
        if err := producer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()

    msg := &sarama.ProducerMessage{Topic: "user", Value: sarama.StringEncoder("this is my message"), Key: sarama.StringEncoder("userid")}
    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        log.Printf("FAILED to send message: %s\n", err)
    } else {
        log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
    }
}