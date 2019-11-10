package dlogger

import (
	"dragon/core/dragon/conf"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

// write log
func writeLog(data interface{}, level string) {
	// 根据data类型删除json或者字符串
	now := time.Now()
	datetime := now.Format("2006-01-02 15:04:05")
	date := now.Format("2006-01-02")
	// 生成或打开文件
	logDir := conf.Conf.Log.Dir
	path := logDir + "/" + date + "." + level + ".log"
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	var logInfo string
	if reflect.TypeOf(data).String() == "string" {
		logInfo = data.(string)
	} else {
		d, _ := json.Marshal(data)
		logInfo = string(d)
	}
	// todo check if safe
	go func() {
		fmt.Fprintf(logFile, "[%s] [%s] || %s \r\n\r\n", datetime, level, logInfo)
		logFile.Close()
	}()
}

func Debug(data interface{}) {
	writeLog(data, "debug")
}

func Info(data interface{}) {
	writeLog(data, "info")
}

func Warn(data interface{}) {
	writeLog(data, "warn")
}

func Error(data interface{}) {
	writeLog(data, "error")
}
