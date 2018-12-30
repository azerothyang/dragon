package dragon

import (
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
	log.Println("start server on " + Conf.Server.Host + ":" + Conf.Server.Port)
	log.Fatal(http.ListenAndServe(Conf.Server.Host+":"+Conf.Server.Port, dragon.Handler))
}
