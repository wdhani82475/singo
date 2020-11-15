package main

import (
	"singo/conf"
	"singo/route"
)

func main() {
	//
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := route.NewRouter()
	r.Run(":3000")
}
