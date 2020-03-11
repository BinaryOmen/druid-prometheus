package utils

import (
	models "druid-prometheus/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPGetMetric returns metric value
func HTTPGetMetric(url string) models.Metric {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("NewRequest construct error : %d", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("GET request error on URL specified : %d", err)

	}

	log.Printf("Response:[%s],Method:[%s]", resp.Status, resp.Request.Method)

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Reading Body Error : %d", err)
	}
	var generic models.Metric

	if err := json.Unmarshal(respData, &generic); err != nil {
		log.Fatalf("Error in unmarshaling request %s", err)
	}

	return generic
}

func url(overlordsep, path string) string {
	url := overlordsep + path
	return url
}
