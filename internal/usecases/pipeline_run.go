package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ti-mo-k/PipelineWatch/internal/domain"
	"github.com/Ti-mo-k/PipelineWatch/internal/repository"
)

type PipelineRunUsecase struct {
	PipelineRepo *repository.PipelineRunRepository
}

func NewPipelineRunUsecase(
	repo *repository.PipelineRunRepository,
) *PipelineRunUsecase {
	return &PipelineRunUsecase{
		PipelineRepo: repo,
	}
}

var ErrInvalidStatus = errors.New("invalid pipeline run status")

func (u *PipelineRunUsecase) CreatePipelineRun(
	ctx context.Context,
	req domain.CreatePipelineRunRequest,
) (domain.PipelineRun, error) {
	return u.PipelineRepo.Create(ctx, req)
}

func (u *PipelineRunUsecase) ListPipelineRuns(ctx context.Context, filter domain.ListPipelineRunsFilter) ([]domain.PipelineRun, error) {
	if filter.Limit <= 0 {
		filter.Limit = 20
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	if filter.Offset < 0 {
		filter.Offset = 0
	}

	if filter.Status != "" {
		switch filter.Status {
		case "queued", "in_progress", "completed":
		default:
			return nil, fmt.Errorf("%w: %s", ErrInvalidStatus, filter.Status)
		}
	}
	return u.PipelineRepo.List(ctx, filter)
}
