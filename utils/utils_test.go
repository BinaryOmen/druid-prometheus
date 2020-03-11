package utils

import (
	"regexp"
	"testing"
)

func TestHTTPGetMetric(t *testing.T) {
	a := HTTPGetMetric("http://localhost:8081/druid/indexer/v1/runningTasks")

	t.Logf("%v", len(a))

}
func WordCount(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S] +`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}
