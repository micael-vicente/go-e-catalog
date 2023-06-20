package models

type ElasticPagedResponse struct {
	Hits ElasticHits `json:"hits"`
}

type ElasticHits struct {
	Total ElasticTotal    `json:"total"`
	Docs  []ElasticSource `json:"hits"`
}

type ElasticSource struct {
	Source Product `json:"_source"`
}

type ElasticTotal struct {
	Value int `json:"value"`
}
