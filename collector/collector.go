package collector

import (
	models "druid-prometheus/model"
	"druid-prometheus/utils"
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	pendingTask   = "/druid/indexer/v1/pendingTasks"
	runningTask   = "/druid/indexer/v1/runningTasks"
	waitingTask   = "/druid/indexer/v1/waitingTasks"
	completedTask = "/druid/indexer/v1/completeTasks"
	datasourceAll = "/druid/coordinator/v1/metadata/datasources?full"
)

var (
	// overlords endpoint.
	overlords = os.Getenv("DRUID_EP")
)

// metricCollector includes the list of metrics
type MetricCollector struct {
	RunningTaskMetric        *prometheus.Desc
	CompletedTaskMetric      *prometheus.Desc
	WaitingTaskMetric        *prometheus.Desc
	PendingTaskMetric        *prometheus.Desc
	SuccessTaskMetric        *prometheus.Desc
	DataSourceCountAllMetric *prometheus.Desc
}

// Collector return the defined metrics with prometheus description
func Collector() *MetricCollector {
	return &MetricCollector{
		RunningTaskMetric: prometheus.NewDesc("druid_running_tasks",
			"Shows number of running tasks",
			nil, prometheus.Labels{
				"tasks": "running",
			},
		),
		CompletedTaskMetric: prometheus.NewDesc("druid_completed_tasks",
			"Shows number of Completed tasks",
			nil, prometheus.Labels{
				"tasks": "completed",
			},
		),
		WaitingTaskMetric: prometheus.NewDesc("druid_waiting_tasks",
			"number of Completed tasks",
			nil, prometheus.Labels{
				"tasks": "waiting",
			},
		),
		PendingTaskMetric: prometheus.NewDesc("druid_pending_tasks",
			"number of Completed tasks",
			nil, prometheus.Labels{
				"tasks": "pending",
			},
		),
		SuccessTaskMetric: prometheus.NewDesc("druid_success_tasks",
			"number of succeeded tasks",
			nil, prometheus.Labels{
				"tasks": "success",
			},
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

	ch <- collector.RunningTaskMetric
	ch <- collector.CompletedTaskMetric
	ch <- collector.PendingTaskMetric
	ch <- collector.WaitingTaskMetric
	ch <- collector.DataSourceCountAllMetric
	ch <- collector.SuccessTaskMetric

}

// Collect ingests
func (collector *MetricCollector) Collect(ch chan<- prometheus.Metric) {

	urlR := overlords + runningTask

	running := utils.HTTPGetMetric(urlR)
	runningTasks := len(running)
	log.Printf("Number of running tasks: %v", runningTasks)
	ch <- prometheus.MustNewConstMetric(collector.RunningTaskMetric, prometheus.GaugeValue, float64(runningTasks))

	urlC := overlords + completedTask
	completed := utils.HTTPGetMetric(urlC)
	completedTasks := len(completed)
	log.Printf("Number of completed tasks: %v", completedTasks)
	ch <- prometheus.MustNewConstMetric(collector.CompletedTaskMetric, prometheus.CounterValue, float64(completedTasks))

	successTask := Status(completed)
	ch <- prometheus.MustNewConstMetric(collector.SuccessTaskMetric, prometheus.CounterValue, float64(successTask))

	urlW := overlords + waitingTask
	waiting := utils.HTTPGetMetric(urlW)
	waitingTasks := len(waiting)
	log.Printf("Number of waiting tasks: %v", waitingTasks)
	ch <- prometheus.MustNewConstMetric(collector.WaitingTaskMetric, prometheus.GaugeValue, float64(waitingTasks))

	urlP := overlords + pendingTask
	pending := utils.HTTPGetMetric(urlP)
	pendingTasks := len(pending)
	log.Printf("Number of pending tasks: %v", pendingTasks)
	ch <- prometheus.MustNewConstMetric(collector.PendingTaskMetric, prometheus.GaugeValue, float64(pendingTasks))

	urlCA := overlords + datasourceAll
	datasourcecountAll := utils.HTTPGetMetric(urlCA)
	datasourcecountAllMetric := len(datasourcecountAll)
	log.Printf("Number of all datasource: %v", datasourcecountAllMetric)
	ch <- prometheus.MustNewConstMetric(collector.DataSourceCountAllMetric, prometheus.CounterValue, float64(datasourcecountAllMetric))

}

func Status(metric models.Metric) (taskCount int) {

	for i := range metric {
		l := models.Label{
			StatusCode: metric[i].StatusCode,
		}
		if l.StatusCode == "SUCCESS" {
			succededTasks := len(metric)
			log.Printf("Number of succeded tasks: %v", succededTasks)
			return succededTasks
		}

	}
	return taskCount
}
