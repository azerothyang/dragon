package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"core/dragon/dlogger"
	"ctrl"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type notFoundHandler struct {
}

var (
	Routes   *httprouter.Router
	testCtrl = &ctrl.Test{} //test controller
)

func init() {
	Routes = httprouter.New()
	Routes.GET("/", testCtrl.Test)
	Routes.GET("/t2", testCtrl.Test2)
	Routes.GET("/t3", testCtrl.Test3)
	Routes.POST("/upload", testCtrl.Upload)
	Routes.GET("/db", testCtrl.GetDBData)
	Routes.GET("/redis", testCtrl.GetRedis)
	Routes.NotFound = notFoundHandler{}
	Routes.PanicHandler = panicHandler
}

// not found route handle
func (notFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>Not Found</h2>")
	//baseCtrl.Json("not found", w)
}

// all panic handler
func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	dlogger.SugarLogger.Errorf("err: %v", err)
	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "<h2>500 Internal Server Error</h2>")
}
