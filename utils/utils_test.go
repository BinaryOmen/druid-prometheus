package utils

import (
	models "druid-prometheus/model"
	"regexp"
	"testing"
)

func TestHTTPGetMetric(t *testing.T) {
	a := HTTPGetMetric("http://localhost:8888/druid/indexer/v1/completeTasks")

	t.Logf("%v", a)
	for i := range a {
		v := models.Label{
			NameDataSource: a[i].NameDataSource,
			StatusCode:     a[i].StatusCode,
		}
		if v.StatusCode == "FAILED" {
			t.Logf("FAILED CODE %s", v.StatusCode)
		}
		t.Logf("%v", v.NameDataSource)
		t.Logf("%v", v.StatusCode)
	}

}
func WordCount(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S] +`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}
