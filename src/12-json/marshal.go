/*
数据的Marshal
*/
package main

import (
	"encoding/json"
	"os"
)

type user struct {
	Name string
	Age  int
}

func main() {
	var auser user
	auser.Name = "lihz"
	auser.Age = 21

	b, err := json.Marshal(auser)
	if err != nil {
		panic(err.Error())
	}
	os.Stdout.Write(b)
}
