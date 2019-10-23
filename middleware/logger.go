package middleware

import (
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
		body, _ := ioutil.ReadAll(r.Body)
		trackMan := &tracker.Tracker{
			SpanId:    spanId,
			Uri:       r.RequestURI,
			Method:    r.Method,
			Header:    r.Header,
			Body:      string(body),
			StartTime: start,
			CostTime:  "",
		}
		trackInfo := trackMan.Marshal()
		w.Header().Set(tracker.TrackKey, trackInfo)
		next.ServeHTTP(w, r)
	})
}
