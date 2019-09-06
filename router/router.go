package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"dragon/core/dragon/dlogger"
	"dragon/ctrl"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type notFoundHandler struct {
}

var (
	Routes      *httprouter.Router
	productCtrl = &ctrl.Product{}   //product controller
)

func init() {
	Routes = httprouter.New()
	Routes.NotFound = notFoundHandler{}
	Routes.PanicHandler = panicHandler
	// -----------------------------商品相关-----------------------------
	// 新增商品
	Routes.POST("/api/product", productCtrl.Add)
	// 伪删除单个商品
	Routes.DELETE("/api/product/:product_code", productCtrl.Delete)
	// 查询商品列表
	Routes.GET("/api/product", productCtrl.GetList)
	// 更新商品信息
	Routes.PUT("/api/product/:product_code", productCtrl.Update)
	// 获取单个商品详情
	Routes.GET("/api/product/:product_code", productCtrl.GetOne)
	// -----------------------------商品相关-----------------------------
}

// not found route handle
func (notFoundHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "text/html; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	fmt.Fprintf(resp, "<h2>Dragon Not Found</h2>")
	//baseCtrl.Json("not found", w)
}

// all panic handler
func panicHandler(resp http.ResponseWriter, req *http.Request, err interface{}) {
	dlogger.SugarLogger.Errorf("err: %v", err)
	resp.Header().Set("content-type", "text/html; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(resp, "<h2>500 Internal Server Error</h2>")
}
