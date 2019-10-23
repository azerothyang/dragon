package ctrl

import (
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/tracker"
	"dragon/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	//品牌模型
	BrandModel = &model.BrandModel{
		BaseModel: model.BaseModel{TableName: model.TBrand{}.TableName()},
	}
	//商品模型
	ProductModel = &model.ProductModel{
		BaseModel: model.BaseModel{TableName: model.TProduct{}.TableName()}, // 传入表名
	}
)

// output struct
type Output struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	SpanId string      `json:"span_id"`
}

type Ctrl struct {
}

func init() {

}

//return with json
func (*Ctrl) Json(data *Output, resp http.ResponseWriter) {
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Content-Length, Accept-Encoding, Origin")

	trackInfo := resp.Header().Get(tracker.TrackKey)
	resp.Header().Del(tracker.TrackKey) // 清除Header中的track
	trackMan := tracker.UnMarshal(trackInfo)
	defer dlogger.Info(trackMan) // 最后写日志跟踪
	trackMan.Resp.Header = resp.Header()
	data.SpanId = trackMan.SpanId
	js, err := json.Marshal(data)
	// 生成耗时
	trackMan.CostTime = time.Since(trackMan.StartTime).String()

	if err != nil {
		trackMan.Error = err
		fmt.Fprint(resp, "error")
		return
	}
	// trackMan data log
	trackMan.Resp.Data = string(js)

	// output
	_, err = resp.Write(js)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		trackMan.Resp.Header = resp.Header()
		trackMan.Error = err
		fmt.Fprint(resp, "")
		return
	}
}
