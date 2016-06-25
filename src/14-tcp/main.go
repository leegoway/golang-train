package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("Hello world\n", time.Now(), "\n"))
		conn.Close()
	}
}
