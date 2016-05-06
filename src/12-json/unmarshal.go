/*
数据的UnMarshal
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	amap := make(map[string]interface{})
	amap["name"] = "贫而乐道"
	amap["url"] = "http://lihuaizhi.com"

	data, err := json.Marshal(amap)
	if err != nil {
		panic(err)
	}

	output := make(map[string]interface{})
	err = json.Unmarshal(data, &output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", output)
}
