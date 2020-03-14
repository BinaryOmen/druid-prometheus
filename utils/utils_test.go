package utils

import (
	"regexp"
	"testing"
)

func TestHTTPGetMetric(t *testing.T) {
	a := HTTPHealthGet("http://localhost:8888/status/health")

	t.Logf("%v", a)

}
func WordCount(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S] +`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}
