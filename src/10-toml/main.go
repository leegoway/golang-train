/*
Toml文件解析
*/
package main

import (
	"10-toml/config"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	configPath = flag.String("config", "config.toml", "config file path")
)

func main() {
	var cfg config.Config
	if _, err := toml.DecodeFile(*configPath, &cfg); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("数据：")
	fmt.Println(cfg)
}
