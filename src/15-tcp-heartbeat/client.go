package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	name = flag.String("name", "defaultClient", "client name to distinguish connect")
	wg   sync.WaitGroup
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", ":9000")
	checkError(err)
	defer conn.Close()

	wg.Add(1)
	go func(c net.Conn) {
		heartbeat := time.Tick(time.Duration(2) * time.Second)
		for {
			select {
			case <-heartbeat:
				fmt.Println("send tick data")
				fmt.Fprintf(conn, "from %s \n", *name)
			}
		}
		wg.Done()
	}(conn)
	wg.Wait()
}
