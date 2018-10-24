package dragon

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Dragon struct {
	Router *httprouter.Router
}

//new一个dragon
func New() *Dragon {
	return new(Dragon)
}

//初始化路由
func (dragon *Dragon)InitRoute(router *httprouter.Router)  {
	dragon.Router = router
}


//开始监听
func (dragon *Dragon)Fly()  {

	//初始化框架
	log.Println("start server on " + Conf.Server.Host + ":" + Conf.Server.Port)
	log.Fatal(http.ListenAndServe(Conf.Server.Host + ":" + Conf.Server.Port, dragon.Router))
}