package repository

import (
	"context"
	"database/sql"

	"github.com/Ti-mo-k/PipelineWatch/internal/domain"
)

type PipelineRunRepository struct {
	db *sql.DB
}

func NewPipelineRunRepository(db *sql.DB) *PipelineRunRepository {
	return &PipelineRunRepository{
		db: db,
	}
}

func (r *PipelineRunRepository) Create(
	ctx context.Context,
	req domain.CreatePipelineRunRequest,
) (domain.PipelineRun, error) {

	query := `
    INSERT INTO pipeline_runs (
        repository,
        workflow_name,
        branch,
        commit_sha,
        status,
        conclusion
    )
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING
        id,
        repository,
        workflow_name,
        branch,
        commit_sha,
        status,
        conclusion,
        started_at,
        finished_at,
        created_at
	`

	var run domain.PipelineRun
	err := r.db.QueryRowContext(
		ctx,
		query,
		req.Repository,
		req.WorkflowName,
		req.Branch,
		req.CommitSHA,
		req.Status,
		req.Conclusion,
	).Scan(
		&run.ID,
		&run.Repository,
		&run.WorkflowName,
		&run.Branch,
		&run.CommitSHA,
		&run.Status,
		&run.Conclusion,
		&run.StartedAt,
		&run.FinishedAt,
		&run.CreatedAt,
	)

	if err != nil {
		return domain.PipelineRun{}, err
	}

	return run, nil
}

func (r *PipelineRunRepository) List(ctx context.Context) ([]domain.PipelineRun, error) {
	query := `
		SELECT
			id,
			repository,
			workflow_name,
			branch,
			commit_sha,
			status,
			conclusion,
			started_at,
			finished_at,
			created_at
		FROM pipeline_runs
		ORDER BY created_at DESC;
	`

	runs := []domain.PipelineRun{}

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var run domain.PipelineRun

		err := rows.Scan(
			&run.ID,
			&run.Repository,
			&run.WorkflowName,
			&run.Branch,
			&run.CommitSHA,
			&run.Status,
			&run.Conclusion,
			&run.StartedAt,
			&run.FinishedAt,
			&run.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		runs = append(runs, run)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return runs, nil
}
