package service

import (
	"context"
	domain "presentation/internal/domain/application"
)

type applicationService struct {
	repository domain.ApplicationRepository
}

func NewApplicationService(r domain.ApplicationRepository) domain.ApplicationService {
	return &applicationService{repository: r}
}

func (a applicationService) Create(ctx context.Context, app domain.Application) error {
	//
	//
	//
	//
	//
	return a.repository.Add(ctx, app)
}

func (a applicationService) Update(ctx context.Context, app domain.Application) error {
	return a.repository.Update(ctx, app)
}

func (a applicationService) Get(ctx context.Context, id string) (domain.Application, error) {

	return a.repository.Get(ctx, id)
}

func (a applicationService) GetAll(ctx context.Context) ([]domain.Application, error) {
	return a.repository.GetAll(ctx)
}

func (a applicationService) Delete(ctx context.Context, id string) error {
	return a.repository.Delete(ctx, id)
}
