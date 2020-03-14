package collector

import (
	"druid-prometheus/utils"
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	druidHealth   = "/status/health"
	pendingTask   = "/druid/indexer/v1/pendingTasks"
	runningTask   = "/druid/indexer/v1/runningTasks"
	waitingTask   = "/druid/indexer/v1/waitingTasks"
	completedTask = "/druid/indexer/v1/completeTasks"
	datasourceAll = "/druid/coordinator/v1/metadata/datasources?full"
)

var (
	// overlords endpoint.
	druid_ep = os.Getenv("DRUID_EP")
)

// metricCollector includes the list of metrics
type MetricCollector struct {
	DruidHealthStatus        *prometheus.Desc
	TaskMetric               *prometheus.Desc
	DataSourceCountAllMetric *prometheus.Desc
}

// Collector return the defined metrics with prometheus description
func Collector() *MetricCollector {
	return &MetricCollector{
		DruidHealthStatus: prometheus.NewDesc("druid_health_status",
			"Health of Druid, 1 is healthy 0 is not",
			nil, prometheus.Labels{
				"druid": "health",
			},
		),
		TaskMetric: prometheus.NewDesc("druid_tasks",
			"Shows number of Druid tasks",
			[]string{"tasks"}, nil,
		),
		DataSourceCountAllMetric: prometheus.NewDesc("druid_datasources_count_all",
			"Returns a list of the names of data sources, regardless of whether there are used segments belonging to those data sources in the cluster or not",
			nil, prometheus.Labels{
				"datasources": "all",
			},
		),
	}
}

// Describe method shall ingest the metric value passed.
func (collector *MetricCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.DruidHealthStatus
	//	ch <- collector.RunningTaskMetric
	ch <- collector.TaskMetric
	//	ch <- collector.PendingTaskMetric
	//	ch <- collector.WaitingTaskMetric
	ch <- collector.DataSourceCountAllMetric

}

// Collect ingests
func (collector *MetricCollector) Collect(ch chan<- prometheus.Metric) {

	urlH := druid_ep + druidHealth
	Health := utils.HTTPHealthGet(urlH)
	log.Printf("Druid Health Status: %v", Health)
	ch <- prometheus.MustNewConstMetric(collector.DruidHealthStatus, prometheus.GaugeValue, Health)

	urls := []string{druid_ep + runningTask, druid_ep + completedTask, druid_ep + pendingTask, druid_ep + waitingTask}
	runningTask := utils.HTTPGetMetric(urls[0])
	completedTask := utils.HTTPGetMetric(urls[1])
	pendingTask := utils.HTTPGetMetric(urls[2])
	waitingTask := utils.HTTPGetMetric(urls[3])

	ch <- prometheus.MustNewConstMetric(collector.TaskMetric, prometheus.GaugeValue, float64(len(runningTask)), "running")
	ch <- prometheus.MustNewConstMetric(collector.TaskMetric, prometheus.GaugeValue, float64(len(completedTask)), "completed")
	ch <- prometheus.MustNewConstMetric(collector.TaskMetric, prometheus.GaugeValue, float64(len(pendingTask)), "pending")
	ch <- prometheus.MustNewConstMetric(collector.TaskMetric, prometheus.GaugeValue, float64(len(waitingTask)), "waiting")

	urlCA := druid_ep + datasourceAll
	datasourcecountAll := utils.HTTPGetMetric(urlCA)
	log.Printf("Number of all datasource: %v", len(datasourcecountAll))
	ch <- prometheus.MustNewConstMetric(collector.DataSourceCountAllMetric, prometheus.CounterValue, float64(len(datasourcecountAll)))

}
