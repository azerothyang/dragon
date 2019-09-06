package ctrl

import (
	"dragon/core/dragon/dlogger"
	"dragon/model"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
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
func (*Ctrl) Json(data interface{}, resp http.ResponseWriter, spanId string) {
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Content-Length, Accept-Encoding, Origin")
	output := data.(Output)
	output.SpanId = spanId
	js, err := json.Marshal(output)
	if err != nil {
		dlogger.SugarLogger.Errorw("Response Marshal Error",
			"Time", time.Now().Format("2006-01-02 15:04:05"),
			"SpanId", spanId,
			"errorInfo", err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}

	// print response data
	dlogger.SugarLogger.Infow("Response Info",
		"Time", time.Now().Format("2006-01-02 15:04:05"),
		"SpanId", spanId,
		"Data", string(js),
	)

	_, err = resp.Write(js)
	if err != nil {
		dlogger.SugarLogger.Errorw("Response Write Error",
			"Time", time.Now().Format("2006-01-02 15:04:05"),
			"SpanId", spanId,
			"errorInfo", err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}
}

// return with msgpack
func (*Ctrl) MsgPack(data interface{}, resp http.ResponseWriter, spanId string) {
	resp.Header().Set("content-type", "text/html;charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Content-Length, Accept-Encoding, Origin")
	output := data.(Output)
	output.SpanId = spanId
	msgp, err := msgpack.Marshal(output)
	if err != nil {
		dlogger.SugarLogger.Errorw("Response Marshal Error",
			"Time", time.Now().Format("2006-01-02 15:04:05"),
			"SpanId", spanId,
			"errorInfo", err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}

	// print response data
	dlogger.SugarLogger.Infow("Response Info",
		"SpanId", spanId,
		"Time", time.Now().Format("2006-01-02 15:04:05"),
		"Data", output,
	)

	_, err = resp.Write(msgp)
	if err != nil {
		dlogger.SugarLogger.Errorw("Response Write Error",
			"Time", time.Now().Format("2006-01-02 15:04:05"),
			"SpanId", spanId,
			"errorInfo", err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}
}
