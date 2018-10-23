package main

import (
	"core/keep"
	"router"
)



func main() {
	//初始化配置
	keep.InitConf()

	//初始化框架
	keeper := keep.New()
	//初始化路由
	keeper.InitRoute(router.Routes)

	//根据配置初始化服务和中间件
	keeper.Run()

}
