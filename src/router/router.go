package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var(
	Routes *httprouter.Router
)

type Info struct {
	Name string
	Age int32
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func init()  {
	Routes = httprouter.New()
	Routes.GET("/", Index)
	Routes.POST("/", Index)
}
