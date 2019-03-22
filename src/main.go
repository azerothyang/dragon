package main

import (
	"core/dragon"
	"core/dragon/conf"
	"core/dragon/dredis"
	"core/dragon/trace"
	"middleware"
	"model"
	"router"
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
