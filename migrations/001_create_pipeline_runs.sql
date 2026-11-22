
CREATE TABLE pipeline_runs (
    id BIGSERIAL PRIMARY KEY,
    repository VARCHAR(255) NOT NULL,
    workflow_name VARCHAR(255) NOT NULL,
    branch VARCHAR(255) NOT NULL,
    commit_sha VARCHAR(40) NOT NULL,

    status VARCHAR(30) NOT NULL,
    conclusion VARCHAR(30),

    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT valid_status
        CHECK (status IN ('queued', 'in_progress', 'completed')),

    CONSTRAINT valid_conclusion
        CHECK (
            conclusion IS NULL OR
            conclusion IN (
                'success',
                'failure',
                'cancelled',
                'timed_out',
                'action_required',
                'skipped',
                'neutral'
            )
        )
);

CREATE INDEX idx_pipeline_runs_repository
ON pipeline_runs(repository);

CREATE INDEX idx_pipeline_runs_status
ON pipeline_runs(status);

CREATE INDEX idx_pipeline_runs_created_at
ON pipeline_runs(created_at DESC);