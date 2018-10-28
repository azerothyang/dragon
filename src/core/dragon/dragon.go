package dragon

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Dragon struct {
	Router *httprouter.Router
}

//new dragon
func New() *Dragon {
	return new(Dragon)
}

// init route
func (dragon *Dragon) InitRoute(router *httprouter.Router) {
	dragon.Router = router
}

//start listening
func (dragon *Dragon) Fly() {

	//dragon fly
	log.Println("start server on " + Conf.Server.Host + ":" + Conf.Server.Port)
	log.Fatal(http.ListenAndServe(Conf.Server.Host+":"+Conf.Server.Port, dragon.Router))
}
