package main

import (
	"core/dragon"
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

	//dragon fly
	dr.Fly()

}
