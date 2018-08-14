package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})
var hdFailures = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	},
	[]string{"device"},
)

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}
