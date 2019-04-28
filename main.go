package main

import (
	"dragon/core/dragon"
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dredis"
	"dragon/core/dragon/trace"
	"dragon/middleware"
	"dragon/model"
	"dragon/router"
)

func main() {
	//init config
	conf.InitConf()

	//init trace (use zipkin)
	trace.Init()

	//init dragon
	dr := dragon.New()

	//init route, you can chain any middleware here :)
	dr.InitRoute(middleware.LogInfo(router.Routes))

	//init db
	model.InitDB()

	//init redis
	dredis.InitRedis()

	//dragon fly
	dr.Fly()

}
