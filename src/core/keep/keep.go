package keep

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type KeepS struct {
	Router *httprouter.Router
}

//new一个keeper
func New() *KeepS {
	return new(KeepS)
}

//初始化路由
func (keeper *KeepS)InitRoute(router *httprouter.Router)  {
	keeper.Router = router
}


//开始监听
func (keeper *KeepS)Run()  {

	//初始化框架
	log.Println("start server on " + Conf.Server.Host + ":" + Conf.Server.Port)
	log.Fatal(http.ListenAndServe(Conf.Server.Host + ":" + Conf.Server.Port, keeper.Router))
}