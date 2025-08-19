package works_chedules

import (
	"reflect"

	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkScheduleService struct {
	repo   *WorkScheduleRepository
	logger *logger.Logger
}

func NewService(module core.Module) core.Provider {
	repo := core.Inject[WorkScheduleRepository](module)
	logger := logger.InjectLog(module)

	return module.NewProvider(&WorkScheduleService{
		repo:   repo,
		logger: logger,
	})
}

func (s *WorkScheduleService) Create(ctx middleware.ContextInfo, model *WorkScheduleModel) (*mongo.InsertOneResult, error) {
	if reflect.ValueOf(model.TenantID).IsZero() {
		model.TenantID = ctx.TenantID
	}

	if reflect.ValueOf(model.CompanyID).IsZero() {
		model.CompanyID = ctx.CompanyID.String()
	}
	data, err := s.repo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create work schedule", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}
	return data, nil
}

func (s *WorkScheduleService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*WorkScheduleModel, error) {
	models, err := s.repo.FindAll(nil, queryParams)
	if err != nil {
		s.logger.Error("Failed to list work schedules", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}
	return models, nil
}
