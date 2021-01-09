package dragon

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/dnacos"
	"dragon/core/dragon/dragonzipkin"
	"dragon/core/dragon/dredis"
	"dragon/repository"
	"dragon/tools/dmongo"
	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
	"log"
)

// AppInit func
// do some components init
func AppInit() {
	//init config
	conf.InitConf()

	// init pprof
	// check if pprof is enabled, then listen port
	if conf.Conf.Server.Pprof.Enabled {
		var host string
		if conf.Conf.Server.Pprof.Host != "" {
			host = conf.Conf.Server.Pprof.Host
		} else {
			host = "0.0.0.0"
		}
		viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeMacarons), viewer.WithAddr(host+":"+conf.Conf.Server.Pprof.Port))
		go func() {
			log.Println("StatsView Pprof server on "+host+":"+conf.Conf.Server.Pprof.Port, "http://"+host+":"+conf.Conf.Server.Pprof.Port+"/debug/statsview")
			mgr := statsview.New()
			defer mgr.Stop()
			// Start() runs a HTTP server at `localhost:18066` by default.
			mgr.Start()
			// Stop() will shutdown the http server gracefully
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
