package main

import (
	"core/dragon"
	"core/dragon/dredis"
	"model"
	"router"
)

func main() {
	//init config
	dragon.InitConf()

	//init dragon
	dr := dragon.New()

	//init route
	dr.InitRoute(router.Routes)

	//init db
	model.InitDB()

	//init redis
	dredis.InitRedis()

	//dragon fly
	dr.Fly()

}
