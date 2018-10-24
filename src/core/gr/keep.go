package gr

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Keeper struct {
	Router *httprouter.Router
}

//new一个keeper
func New() *Keeper {
	return new(Keeper)
}

//初始化路由
func (keeper *Keeper)InitRoute(router *httprouter.Router)  {
	keeper.Router = router
}


//开始监听
func (keeper *Keeper)Run()  {

	//初始化框架
	log.Println("start server on " + Conf.Server.Host + ":" + Conf.Server.Port)
	log.Fatal(http.ListenAndServe(Conf.Server.Host + ":" + Conf.Server.Port, keeper.Router))
}