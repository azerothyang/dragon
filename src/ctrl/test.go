package ctrl

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Test struct {
	Ctrl
}

func (t *Test)Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	t.Json("hello world", w)
}