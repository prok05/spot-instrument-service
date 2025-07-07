package metric

import (
	"github.com/prok05/spot-instrument-service/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Start(cfg *config.Config) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(cfg.Prometheus.Addr, nil)
}
