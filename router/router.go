package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/tracker"
	"dragon/ctrl"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type notFoundHandler struct {
}

var (
	Routes      *httprouter.Router
	productCtrl = &ctrl.Product{} //product controller
)

func init() {
	Routes = httprouter.New()
	Routes.NotFound = notFoundHandler{}
	Routes.PanicHandler = panicHandler
	// -----------------------------商品相关-----------------------------
	Routes.GET("/test", productCtrl.Test)
	// 新增商品
	Routes.POST("/api/product", productCtrl.Add)
	// -----------------------------商品相关-----------------------------
}

// not found route handle
func (notFoundHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "text/html; charset=utf-8")
	resp.Header().Set("x-server", "dragon")

	trackInfo := resp.Header().Get(tracker.TrackKey)
	resp.Header().Del(tracker.TrackKey) // 清除Header中的track
	trackMan := tracker.UnMarshal(trackInfo)
	trackMan.Resp.Header = resp.Header()
	trackMan.Resp.Data = "<h2>Dragon Not Found</h2>"
	trackMan.CostTime = time.Since(trackMan.StartTime).String()
	dlogger.Info(trackMan) // 最后写日志跟踪
	resp.Write([]byte("<h2>Dragon Not Found</h2>"))
}

// all panic handler
func panicHandler(resp http.ResponseWriter, req *http.Request, err interface{}) {
	resp.Header().Set("content-type", "text/html; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.WriteHeader(http.StatusInternalServerError)

	trackInfo := resp.Header().Get(tracker.TrackKey)
	resp.Header().Del(tracker.TrackKey) // 清除Header中的track
	trackMan := tracker.UnMarshal(trackInfo)
	trackMan.Resp.Header = resp.Header()
	trackMan.Resp.Data = "<h2>500 Internal Server Error</h2>"
	trackMan.Error = err
	dlogger.Error(trackMan) // 写入日志跟踪
	resp.Write([]byte("<h2>500 Internal Server Error</h2>"))
}
