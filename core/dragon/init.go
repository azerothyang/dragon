package dragon

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/dnacos"
	"dragon/core/dragon/dragonzipkin"
	"dragon/core/dragon/dredis"
	"dragon/domain/repository"
	"dragon/tools/dmongo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

// AppInit func
// do some components init
func AppInit() {
	//init config
	conf.InitConf()

	// init Prometheus pprof
	if conf.Conf.Server.Pprof.Enabled {
		var host string
		if conf.Conf.Server.Pprof.Host != "" {
			host = conf.Conf.Server.Pprof.Host
		} else {
			host = "0.0.0.0"
		}
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			log.Println("Prometheus Pprof server on "+host+":"+conf.Conf.Server.Pprof.Port, "http://"+host+":"+conf.Conf.Server.Pprof.Port+"/metrics")
			http.ListenAndServe(host+":"+conf.Conf.Server.Pprof.Port, nil)
		}()
	}

	// init zipkin server middleware
	if conf.Conf.Zipkin.Enable {
		dragonzipkin.Init()
	}
	if conf.Conf.Nacos.Enable {
		dnacos.Init()
	}

	//init db
	repository.InitDB()

	//init redis
	dredis.InitRedis()

	// init mongodb
	dmongo.InitDB()

	// init logger
	go func() {
		dlogger.TickFlush()
	}()
}
