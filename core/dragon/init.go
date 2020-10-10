package dragon

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dredis"
	"dragon/repository"
	"dragon/tools/dmongo"
	"log"
	"net/http"
)

// AppInit func
// do some components init
func AppInit() {
	//init config
	conf.InitConf()

	// init pprof
	// check if pprof is enabled, then listen port
	if conf.Conf.Server.Pprof.Enabled {
		go func() {
			err := http.ListenAndServe(conf.Conf.Server.Pprof.Host+":"+conf.Conf.Server.Pprof.Port, nil)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	//init db
	repository.InitDB()

	//init redis
	dredis.InitRedis()

	// init mongodb
	dmongo.InitDB()
}
