package dto

import "plagChecker/internal/model"

type CheckStudentResponse struct {
	StudentCheckResults []model.StudentCheckResult `json:"student_check_result"`
}

type CheckLabResponse struct {
	LabCheckResults []model.LabCheckResult `json:"lab_check_results"`
}
