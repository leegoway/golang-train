package main 

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"sync"
	"strings"
	"log"
)

var (
    kafka = "127.0.0.1:9092,127.0.0.1:9091,127.0.0.1:9093"
    wg     sync.WaitGroup
    logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger
	// 连接kafka消息服务器
    consumer, err := sarama.NewConsumer(strings.Split(kafka, ","), nil)
    if err != nil {
        logger.Println("Failed to start consumer: %s", err)
    }

    // consumer.Partitions 用户获取Topic上所有的Partitions. 消息服务器上已经创建了test这个topic,所以,在这里指定参数为test.
    partitionList, err := consumer.Partitions("user")
    if err != nil {
        logger.Println("Failed to get the list of partitions: ", err)
    }

    for partition := range partitionList {
        pc, err := consumer.ConsumePartition("user", int32(partition), sarama.OffsetNewest)
        if err != nil {
            logger.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
        }
        defer pc.AsyncClose()

        wg.Add(1)

        go func(sarama.PartitionConsumer) {
            defer wg.Done()
            for msg := range pc.Messages() {
                fmt.Println("message is :", msg)
                fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
                fmt.Println()
            }
        }(pc)
    }

    wg.Wait()

    logger.Println("Done consuming topic hello")
    consumer.Close()
}