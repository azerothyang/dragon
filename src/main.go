package main

import (
	"core/dragon"
	"router"
)



func main() {
	//初始化配置
	dragon.InitConf()

	//初始化框架
	dr := dragon.New()
	//初始化路由
	dr.InitRoute(router.Routes)

	//根据配置初始化服务和中间件
	dr.Fly()

}
