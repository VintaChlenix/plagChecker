package model

type Metadata struct {
	Name     string `json:"name"`
	LabID    int    `json:"lab_id"`
	Variant  int    `json:"variant"`
	NormCode string `json:"norm_code"`
	Sum      string `json:"sum"`
}
