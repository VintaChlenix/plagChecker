package model

type Sending struct {
	Name    string    `json:"name"`
	LabID   string    `json:"lab_id"`
	Variant string    `json:"variant"`
	Results []float64 `json:"results"`
}
