package main

import (
	"ttc-go/config"
	"ttc-go/server"
)

func main() {
	// 从配置文件读取配置
	config.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
