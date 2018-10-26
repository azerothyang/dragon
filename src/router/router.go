package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"ctrl"
	"github.com/julienschmidt/httprouter"
)

var(
	Routes *httprouter.Router
)

func init()  {
	Routes = httprouter.New()
	Routes.GET("/", (&ctrl.Test{}).Test)
	Routes.POST("/upload", (&ctrl.Test{}).Upload)
	Routes.GET("/db", (&ctrl.Test{}).GetDBData)
}
