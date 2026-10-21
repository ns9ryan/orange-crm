package main

import (
	"orange-crm/initialize"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	// 设置为 standalone 模式
	if err := os.Setenv("RUN_AS_STANDALONE", "1"); err != nil {
		panic(err)
	}

	initialize.Run()
}
