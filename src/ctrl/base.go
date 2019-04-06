package ctrl

import (
	"core/dragon/dlogger"
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"time"
)

var (
	testModel = &model.TestModel{}
)

type Ctrl struct {
}

func init() {

}

//return with json
func (*Ctrl) Json(data interface{}, resp http.ResponseWriter) {
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("server", "dragon")
	js, err := json.Marshal(data)
	if err != nil {
	    dlogger.SugarLogger.Error(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}

	// print response data
	dlogger.SugarLogger.Infow("Response Data Info:",
		"Data", string(js),
		"Time", time.Now().Format("2006-01-02 15:04:05"),
	)

	_, err = resp.Write(js)
	if err != nil {
		dlogger.SugarLogger.Error(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}
}
