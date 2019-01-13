package ctrl

import (
	"core/dragon"
	"core/dragon/dredis"
	"dto"
	"github.com/julienschmidt/httprouter"
	"model"
	"net/http"
)

type Test struct {
	Ctrl
}

func (t *Test) Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	requests := dragon.Parse(r)
	tt := model.Test{}
	dto.TestPToTestS(requests, &tt)
	t.Json("hello world", w)
}

// upload test
func (t *Test) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dragon.Upload(r, "file", "./test.png")
	t.Json("upload success", w)
}

// mysql test
func (t *Test) GetDBData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//testModel.Create()
	//testModel.Update()
	res := testModel.Get()
	output := dto.TestSToTest(res)
	t.Json(output, w)
}

// redis test
func (t *Test) GetRedis(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := dredis.Redis.Get("x").Result()
	t.Json(res, w)
}
