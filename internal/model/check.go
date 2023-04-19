package model

type StudentCheckResult struct {
	LabID   string `json:"lab_id"`
	Variant string `json:"variant"`
}

type LabCheckResult struct {
	Name      string    `json:"name"`
	Variant   string    `json:"variant"`
	Results   []float64 `json:"results"`
	URL       string    `json:"url"`
	SourceURL string    `json:"source_url"`
}
