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
	testModel = &model.TestModel{}
)

// output struct
type Output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Ctrl struct {
}

func init() {

}

//return with json
func (*Ctrl) Json(data interface{}, resp http.ResponseWriter, spanId string) {
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("server", "dragon")
	js, err := json.Marshal(data)
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
	resp.Header().Set("server", "dragon")
	msgp, err := msgpack.Marshal(data)
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
		"Data", data,
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
