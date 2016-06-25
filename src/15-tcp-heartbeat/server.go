package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Accept connection: ", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	ln, err := net.Listen("tcp", ":9000")
	checkError(err)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handle(conn)
	}

}
