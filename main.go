package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/viggin543/gke-gitops-tutorial-cloudbuild/handlers"
	"net/http"
)

func main() {

	finish := make(chan bool)

	applicationServer := http.NewServeMux()
	metricsServer := http.NewServeMux()

	applicationServer.Handle("/banana", promhttp.InstrumentHandlerCounter(
		promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hello_requests_total",
				Help: "Total number of hello-world requests by HTTP code.",
			},
			[]string{"code"},
		),
		handlers.HelloWorldHandler,
	))
	metricsServer.Handle("/", promhttp.Handler())
	go func() {
		_ = http.ListenAndServe(":9123", metricsServer)
	}()
	go func() {
		_ = http.ListenAndServe(":8080", applicationServer)
	}()
	<-finish
}
