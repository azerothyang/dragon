package dlogger

import "go.uber.org/zap"

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

// zap link: https://github.com/uber-go/zap
func init() {
	Logger, _ = zap.NewProduction()
	SugarLogger = Logger.Sugar()
}
