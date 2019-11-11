package middleware

import (
	"bytes"
	"dragon/core/dragon/tracker"
	"dragon/core/dragon/util"
	"io/ioutil"
	"net/http"
	"time"
)

func LogInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// todo parse params will sometimes happen twice, because in controller will call it

		spanId, _ := util.NewUUID()
		// 读取
		body, _ := ioutil.ReadAll(r.Body)
		// 把刚刚读出来的再写进去
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		trackMan := &tracker.Tracker{
			SpanId:    spanId,
			Uri:       r.RequestURI,
			Method:    r.Method,
			ReqHeader: r.Header,
			Body:      string(body),
			StartTime: start,
			CostTime:  "",
		}
		trackInfo := trackMan.Marshal()
		r.Header.Set(tracker.TrackKey, trackInfo)
		next.ServeHTTP(w, r)
	})
}
