package trace

import (
	"core/dragon/conf"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

// zipkin http client
var Client *zipkinhttp.Client

func Init() {
	// if not enable, return
	if !conf.Conf.Zipkin.Enable {
		return
	}
	reporter := httpreporter.NewReporter(conf.Conf.Zipkin.Url)
	//defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(conf.Conf.Zipkin.ServiceName, conf.Conf.Zipkin.Endpoint)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	// initialize our tracer
	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// create global zipkin traced http client
	Client, err = zipkinhttp.NewClient(tracer, zipkinhttp.ClientTrace(true))
	if err != nil {
		log.Fatalf("unable to create client: %+v\n", err)
	}
}
