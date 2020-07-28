package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//config struct
type ConfS struct {
	Server struct {
		Host  string
		Port  string
		Pprof struct {
			Enabled bool
			Host    string
			Port    string
		}
	}
	Database struct {
		Mysql struct {
			Master struct {
				Host     string
				Port     string
				User     string
				Password string
				Database string
				Charset  string
				Timeout  string
				MaxIdle  int
				MaxConn  int
			}
		}
		Redis struct {
			Host    string
			Port    string
			Auth    string
			Timeout string
			Db      int
		}
		Mongodb struct {
			Host     string
			Port     string
			Username string
			Password string
			Database string
			Timeout  uint8
		}
	}
	Kafka struct {
		Broker string
	}
	Log struct {
		Dir string
	}
}

var (
	Conf    ConfS
	Env     = "dev"
	ExecDir = "" // current exec file path
)

//init config
func InitConf() {
	dir, err := GetCurrentPath()
	ExecDir = dir
	if err != nil {
		log.Fatal(err)
	}

	var envb []byte
	// read DRAGON env first
	env := os.Getenv("DRAGON")
	if env == "" {
		envb, err = ioutil.ReadFile(dir + FmtSlash("conf/.env"))

		// check last char is LF (\n)
		if envb[len(envb)-1] == 10 {
			envb = envb[:len(envb)-1]
		}
	} else {
		envb = []byte(env)
	}
	if err != nil {
		log.Fatal(err)
	}

	Env = string(envb)
	// check Env != dev,prod,test
	if (Env != "dev") && (Env != "test") && (Env != "prod") {
		panic("environment variable DRAGON can only be dev,test or prod")
	}
	var config []byte
	config, err = ioutil.ReadFile(dir + FmtSlash("conf/"+Env+".yml"))
	//check env DRAGON or release/.env
	if err != nil {
		// read yml config fail, return fail
		panic("release/conf/" + Env + ".yml not found")
	}
	err = yaml.Unmarshal(config, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}

//get current exec file path
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
	return path[0 : i+1], nil
}

// according to operating system to change path slash, default use linux path slash
func FmtSlash(path string) string {
	sys := runtime.GOOS
	if sys == `windows` {
		return strings.Replace(path, "/", "\\", -1)
	}
	return path
}
