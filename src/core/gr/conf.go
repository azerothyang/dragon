package gr

import (
	"core/yaml"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var	(
	Conf ConfS
	Env = "debug"
)

type ConfS struct {
	Server struct{
		Host string
		Port string
	}
} 

//初始化配置文件
func InitConf() {
	dir, err := GetCurrentPath()
	if err != nil {
		log.Fatal(err)
	}

	var envb []byte
	envb, err = ioutil.ReadFile(dir + FmtSlash("conf/.env"))
	if err != nil {
		log.Fatal(err)
	}

	var config []byte
	//如果.env不等于debug则是生产环境
	if string(envb) != "debug" {
		Env = "production"
		config, err = ioutil.ReadFile(dir + FmtSlash("conf/prod.yml"))
	} else {
		Env = "debug"
		config, err = ioutil.ReadFile(dir + FmtSlash("conf/dev.yml"))
	}
	err = yaml.Unmarshal(config, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}

//获取当前执行文件路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}


//根据操作系统，如果linux保持不变; 如果windows,将linux路径换为window路径。
func FmtSlash(path string) string  {
	sys := runtime.GOOS
	if sys == `windows`{
		return strings.Replace(path, "/", "\\", -1)
	}
	return path
}