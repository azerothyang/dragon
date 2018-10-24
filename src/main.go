package main

import (
	"core/gr"
	"router"
)



func main() {
	//初始化配置
	gr.InitConf()

	//初始化框架
	keeper := gr.New()
	//初始化路由
	keeper.InitRoute(router.Routes)

	//根据配置初始化服务和中间件
	keeper.Run()

}
