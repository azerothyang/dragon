package middleware

import (
	"core/dragon"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

// zap link: https://github.com/uber-go/zap
func init() {
	Logger, _ = zap.NewProduction()
	SugarLogger = Logger.Sugar()
}

func LogInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// todo parse params will sometimes happen twice, because in controller will call it
		SugarLogger.Infow("Request Info:",
			"Method", r.Method,
			"Path", r.URL.Path,
			"Time", start.Format("2006-01-02 15:04:05"),
			"Params", dragon.Parse(r),
		)
		next.ServeHTTP(w, r)
		SugarLogger.Infow("Request Finish Info:",
			"Method", r.Method,
			"Path", r.URL.Path,
			"Time", start.Format("2006-01-02 15:04:05"),
			"CostTime", time.Since(start).String(),
		)
		Logger.Sync() // flushes buffer, if any
	})
}
