package main

import (
	"ak520hl.cn/ak520hl/douban_rent/douban"
	"ak520hl.cn/ak520hl/douban_rent/worker"
	"fmt"
)

func main() {
	fmt.Println("project start")
	client := douban.NewClient()

	w := worker.New(client)

	w.Run()
}
