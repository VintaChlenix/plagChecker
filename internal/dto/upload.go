package dto

import "plagChecker/internal/model"

type UploadRequest struct {
	model.UploadInfo `json:"upload_info"`
}

type SelectStudentMetadataResponse struct {
	StudentMetadata model.Metadata `json:"student_metadata"`
}

type CountMetadataResponse struct {
	Metadata model.Metadata `json:"metadata"`
}

type CheckMetadataResponse struct {
	Result          model.Result `json:"result"`
	Explanation     string       `json:"explanation"`
	MatchPercentage []float64    `json:"match_ercentage"`
	Source          string       `json:"source"`
	SourceURL       string       `json:"source_url"`
}
