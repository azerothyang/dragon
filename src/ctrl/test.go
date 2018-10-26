package ctrl

import (
	"core/dragon"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Test struct {
	Ctrl
}

func (t *Test)Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	t.Json("hello world", w)
}

func (t *Test)Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	dragon.Upload(r, "file", "./test.png")
	t.Json("upload success", w)
}

func (t *Test)GetDBData(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	res := testModel.Get()
	t.Json(res, w)
}