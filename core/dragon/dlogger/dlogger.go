package dlogger

import (
	"dragon/core/dragon/conf"
	"dragon/tools"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	DebugLevel    = "debug"
	InfoLevel     = "info"
	WarnLevel     = "warn"
	ErrorLevel    = "error"
	SqlInfoLevel  = "sql"
	SqlErrorLevel = "sql.error"
)

// write log
func writeLog(level string, data ...interface{}) {
	if data == nil || len(data) == 0 {
		// 如果data为空，不进行打印
		return
	}
	// 根据data类型删除json或者字符串
	now := time.Now()
	datetime := now.Format("2006-01-02 15:04:05")
	date := now.Format("2006-01-02")
	// 生成或打开文件
	logDir := conf.Conf.Log.Dir
	path := conf.ExecDir + "/" + logDir + "/" + date + "." + level + ".log"
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer logFile.Close()
	if err != nil {
		log.Println(fmt.Sprintf("error:%+v", err))
	}
	var logInfo string
	d, _ := tools.FastJson.Marshal(&data)
	logInfo = string(d)
	// todo check if safe
	fmt.Fprintf(logFile, "[%s] [%s] || %s \r\n\r\n", datetime, level, logInfo)
}

func Debug(data ...interface{}) {
	writeLog(DebugLevel, data...)
}

func Info(data ...interface{}) {
	writeLog(InfoLevel, data...)
}

func Warn(data ...interface{}) {
	writeLog(WarnLevel, data...)
}

func Error(data ...interface{}) {
	writeLog(ErrorLevel, data...)
}

func SqlInfo(data ...interface{}) {
	writeLog(SqlInfoLevel, data...)
}

func SqlError(data ...interface{}) {
	writeLog(SqlErrorLevel, data...)
}
