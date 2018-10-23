package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var(
	Routes *httprouter.Router
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func init()  {
	Routes = httprouter.New()
	Routes.GET("/", Index)
}
