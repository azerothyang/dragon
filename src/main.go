package main

import (
	"core/dragon"
	"core/dragon/dredis"
	"middleware"
	"model"
	"router"
)

func main() {
	//init config
	dragon.InitConf()

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
