package dragon

import (
	"dragon/core/dragon/conf"
	"log"
	"net/http"
)

type Dragon struct {
	Handler http.Handler
}

//new dragon
func New() *Dragon {
	return new(Dragon)
}

// init route
func (dragon *Dragon) InitRoute(handler http.Handler) {
	dragon.Handler = handler
}

//start listening
func (dragon *Dragon) Fly() {

	//dragon fly
	log.Println("env: " + conf.Env)
	log.Println("set environment variable DRAGON dev,test,prod 🐲🐲🐲")
	webAddr := "http://0.0.0.0:" + conf.Conf.Server.Port
	if conf.Conf.Server.Host != ""{
		webAddr = "http://" + conf.Conf.Server.Host + ":" + conf.Conf.Server.Port
	}
	log.Println("start server on: " + webAddr)
	log.Fatal(http.ListenAndServe(conf.Conf.Server.Host+":"+conf.Conf.Server.Port, dragon.Handler))
}
