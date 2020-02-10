package main

import (
	"dragon/core/dragon"
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dredis"
	"dragon/middleware"
	"dragon/model"
	"dragon/router"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//init config
	conf.InitConf()

	// init pprof
	// check if pprof is enabled, then listen port
	if conf.Conf.Server.Pprof.Enabled {
		go func() {
			err := http.ListenAndServe(conf.Conf.Server.Pprof.Host+":"+conf.Conf.Server.Pprof.Port, nil)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}


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
