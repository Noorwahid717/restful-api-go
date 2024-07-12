package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type EventUsecase struct {
	base *BaseUsecase[model.Event, dto.CreateEvent, dto.UpdateEvent, dto.Event]
}

func NewEventUsecase(cfg *config.Config, repository repository.EventRepository) *EventUsecase {
	return &EventUsecase{
		base: NewBaseUsecase[model.Event, dto.CreateEvent, dto.UpdateEvent, dto.Event](cfg, repository),
	}
}

// Create
func (u *EventUsecase) Create(ctx context.Context, req dto.CreateEvent) (dto.Event, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *EventUsecase) Update(ctx context.Context, id int, req dto.UpdateEvent) (dto.Event, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *EventUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *EventUsecase) GetById(ctx context.Context, id int) (dto.Event, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *EventUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Event], error) {
	return s.base.GetByFilter(ctx, req)
}

// // Get list
// func (s *EventUsecase) GetList(ctx context.Context) ([]dto.Event, error) {
// 	return s.base.GetList(ctx)
// }
