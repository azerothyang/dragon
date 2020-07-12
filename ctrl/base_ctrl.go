package ctrl

import (
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/tracker"
	"dragon/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	//商品模型
	ProductModel = model.ProductModel{
		BaseModel: model.BaseModel{TableName: model.TProduct{}.TableName()}, // 传入表名
	}
)

// output struct
type Output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// output data structure
type OutData struct {
	Output
	SpanId string `json:"span_id"`
}

// Controller interface
type Controller interface {
	InitReqAndResp(req *http.Request, resp http.ResponseWriter)
	GetRequestParams() map[string]string
	BindRequestJsonToStruct(data interface{}) error
	Json(data *Output)
}

// Ctrl struct
type Ctrl struct {
	req  *http.Request
	resp http.ResponseWriter
}

// init requests and response bind to ctrl struct
func (ctrl *Ctrl) InitReqAndResp(req *http.Request, resp http.ResponseWriter) {
	ctrl.req = req
	ctrl.resp = resp
}

// get request params (get and post params)
func (ctrl Ctrl) GetRequestParams() map[string]string {
	requests := make(map[string]string)
	var err error
	err = ctrl.req.ParseForm()
	if err != nil {
		log.Println(err)
		return requests
	}

	for k, v := range ctrl.req.Form {
		if len(v) == 1 {
			requests[k] = v[0]
		}
	}

	return requests
}

//parse raw json bind to struct
func (ctrl Ctrl) BindRequestJsonToStruct(data interface{}) error {
	var body []byte
	var err error
	body, err = ioutil.ReadAll(ctrl.req.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// statusCode with httpStatus code
//return with json
func (ctrl Ctrl) Json(data *Output, statusCode int) {
	resp := ctrl.resp
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Content-Length, Accept-Encoding, Origin")

	trackInfo := ctrl.req.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)
	defer func() {
		dlogger.Info(trackMan) // 最后写日志跟踪
	}()
	trackMan.Resp.Header = resp.Header()
	outData := OutData{
		Output: *data,
		SpanId: trackMan.SpanId,
	}
	js, err := json.Marshal(outData)
	// 生成耗时
	trackMan.CostTime = time.Since(trackMan.StartTime).String()

	if err != nil {
		trackMan.Error = err
		fmt.Fprint(resp, "error")
		return
	}
	// trackMan data log
	resp.WriteHeader(statusCode)
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
