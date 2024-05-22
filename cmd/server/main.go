package main

import (
	"fmt"

	"github.com/fabiokusaba/api-go-todolist/configs"
)

func main() {
	conf, _ := configs.LoadConf(".")
	fmt.Println(conf.DatabaseName)
}