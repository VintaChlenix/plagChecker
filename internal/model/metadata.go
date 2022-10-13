package model

type Metadata struct {
	Name     string `json:"name"`
	LabID    string `json:"lab_id"`
	Variant  string `json:"variant"`
	NormCode string `json:"norm_code"`
	Sum      string `json:"sum"`
	Tokens   string `json:"tokens"`
}

type Result string

const (
	CheckResultType1 Result = "Completely plagiarized"
	CheckResultType2 Result = "Renamed variables, functions and so on..."
	CheckResultType3 Result = "A little plagiarized"
	CheckResultType4 Result = "Not Plagiarism"
)
