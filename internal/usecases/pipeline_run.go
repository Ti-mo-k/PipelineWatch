package usecases

import (
	"context"

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

func (u *PipelineRunUsecase) CreatePipelineRun(
	ctx context.Context,
	req domain.CreatePipelineRunRequest,
) (domain.PipelineRun, error) {
	return u.PipelineRepo.Create(ctx, req)
}
