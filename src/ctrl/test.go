package ctrl

import (
	"core/dragon"
	"core/dragon/dredis"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Test struct {
	Ctrl
}

func (t *Test) Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.Json("hello world", w)
}

// upload test
func (t *Test) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dragon.Upload(r, "file", "./test.png")
	t.Json("upload success", w)
}

// mysql test
func (t *Test) GetDBData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := testModel.Get()
	t.Json(res, w)
}

// redis test
func (t *Test)GetRedis(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	res, _ := dredis.Redis.Get("x").Result()
	t.Json(res, w)
}
