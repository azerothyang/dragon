package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"ctrl"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type notFoundHandler struct {
}

var (
	Routes *httprouter.Router
	testCtrl = &ctrl.Test{}  //test controller
)

func init() {
	Routes = httprouter.New()
	Routes.GET("/", testCtrl.Test)
	Routes.POST("/upload", testCtrl.Upload)
	Routes.GET("/db", testCtrl.GetDBData)
	Routes.GET("/redis", testCtrl.GetRedis)
	Routes.NotFound = notFoundHandler{}
}

// not found route handle
func (notFoundHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>Not Found</h2>")
	//baseCtrl.Json("not found", w)
}
