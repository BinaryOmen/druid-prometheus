package models

type Metric []struct {
	NameDataSource string `json:"dataSource"`
	StatusCode     string `json:"statusCode"`
}

type Label struct {
	NameDataSource string `json:"dataSource"`
	StatusCode     string
}
