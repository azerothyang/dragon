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
	log.Println("set environment variable DRAGON debug or production 😄🐲😄")
	log.Println("start server on " + conf.Conf.Server.Host + ":" + conf.Conf.Server.Port)
	log.Fatal(http.ListenAndServe(conf.Conf.Server.Host+":"+conf.Conf.Server.Port, dragon.Handler))
}
