package dto

import "plagChecker/internal/model"

type CheckStudentResponse struct {
	StudentCheckResult []model.StudentCheckResult `json:"student_check_result"`
}
