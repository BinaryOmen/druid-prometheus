package models

type Metric []struct {
	NameDataSource string `json:"name"`
	StatusCode     string `json:"statusCode"`
}
