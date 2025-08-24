package titles

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type TitleService struct {
	titleRepo      *Repository
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
}

func NewService(module core.Module) core.Provider {
	titleRepo := module.Ref(REPOSITORY).(*Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&TitleService{
		titleRepo:      titleRepo,
		logger:         logger,
		eventPublisher: eventPublisher,
	})
}

func (s *TitleService) Create(ctx middleware.ContextInfo, model *TitleModel) (*TitleModel, error) {
	createdTitle, err := s.titleRepo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create title", logger.Metadata{
			"error": err.Error(),
			"input": model,
		})
		return nil, err
	}
	return createdTitle, nil
}

func (s *TitleService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*TitleModel, int64, error) {
	filter := make(map[string]any)
	filter["tenant_id"] = ctx.TenantID

	if ctx.CompanyID != uuid.Nil {
		filter["company_id"] = ctx.CompanyID
	}

	titles, total, err := s.titleRepo.FindAll(filter, sqlorm.FindOptions{
		Limit:  queryParams.Limit,
		Offset: queryParams.Skip,
	})
	if err != nil {
		s.logger.Error("Failed to list titles", logger.Metadata{
			"error":  err.Error(),
			"filter": filter,
		})
		return nil, 0, err
	}
	return titles, total, nil
}

func (s *TitleService) GetByID(ctx middleware.ContextInfo, id string) (*TitleModel, error) {
	title, err := s.titleRepo.FindByID(id)
	if err != nil {
		s.logger.Error("Failed to fetch title", logger.Metadata{
			"error": err.Error(),
			"id":    id,
		})
		return nil, err
	}
	return title, nil
}

func (s *TitleService) UpdateByID(ctx middleware.ContextInfo, id string, model *TitleModel) (*TitleModel, error) {
	updatedTitle, err := s.titleRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to update title", logger.Metadata{
			"error": err.Error(),
			"input": model,
		})
		return nil, err
	}
	return updatedTitle, nil
}

func (s *TitleService) DeleteByID(ctx middleware.ContextInfo, id string) error {
	err := s.titleRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error("Failed to delete title", logger.Metadata{
			"error": err.Error(),
			"id":    id,
		})
		return err
	}
	return nil
}
