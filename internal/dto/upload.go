package dto

import "plagChecker/internal/model"

type UploadRequest struct {
	StudentMetadata model.Metadata `json:"student_metadata"`
}

type SelectStudentMetadataResponse struct {
	StudentMetadata model.Metadata `json:"student_metadata"`
}
