/*
flag包解析参数。例如go run main.go -u mysql -p nopass.2
*/
package main

import (
	"flag"
	"fmt"
)

var (
	username = flag.String("u", "root", "mysql username -u")
	password = flag.String("p", "", "mysql password -p")
)

func main() {
	flag.Parse()
	fmt.Println(*username)
	fmt.Println(*password)
}
