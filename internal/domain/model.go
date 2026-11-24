package domain

import "time"

type CreatePipelineRunRequest struct {
	Repository   string  `json:"repository"`
	WorkflowName string  `json:"workflow_name"`
	Branch       string  `json:"branch"`
	CommitSHA    string  `json:"commit_sha"`
	Status       string  `json:"status"`
	Conclusion   *string `json:"conclusion"`
}

type PipelineRun struct {
	ID           int64      `json:"id"`
	Repository   string     `json:"repository"`
	WorkflowName string     `json:"workflow_name"`
	Branch       string     `json:"branch"`
	CommitSHA    string     `json:"commit_sha"`
	Status       string     `json:"status"`
	Conclusion   *string    `json:"conclusion"`
	StartedAt    *time.Time `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

type ListPipelineRunsFilter struct {
	Status     string
	Repository string
	Limit      int
	Offset     int
}
