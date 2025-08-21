package grades

import (
	"reflect"

	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type GradeService struct {
	gradeRepo      *Repository
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
}

func NewService(module core.Module) core.Provider {
	gradeRepo := module.Ref(REPOSITORY).(*Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&GradeService{
		gradeRepo:      gradeRepo,
		logger:         logger,
		eventPublisher: eventPublisher,
	})
}

func (s *GradeService) Create(ctx middleware.ContextInfo, model *GradeModel) (*GradeModel, error) {
	createdGrade, err := s.gradeRepo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create grade", logger.Metadata{
			"error": err.Error(),
			"input": model,
		})
		return nil, err
	}
	return createdGrade, nil
}

func (s *GradeService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*GradeModel, int64, error) {
	filter := make(map[string]any)
	filter["tenant_id"] = ctx.TenantID

	if !reflect.ValueOf(ctx.CompanyID).IsZero() {
		filter["company_id"] = ctx.CompanyID
	}

	data, total, err := s.gradeRepo.FindAll(filter, sqlorm.FindOptions{
		Limit:  queryParams.Limit,
		Offset: queryParams.Skip,
	})
	if err != nil {
		s.logger.Error("Failed to fetch list grades", logger.Metadata{
			"error":  err.Error(),
			"filter": filter,
		})
		return nil, 0, err
	}
	return data, total, nil
}

func (s *GradeService) UpdateByID(ctx middleware.ContextInfo, id string, model *GradeModel) (*GradeModel, error) {
	updatedGrade, err := s.gradeRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to update grade", logger.Metadata{
			"error": err.Error(),
			"input": model,
		})
		return nil, err
	}
	return updatedGrade, nil
}

func (s *GradeService) DeleteByID(ctx middleware.ContextInfo, id string) error {
	err := s.gradeRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error("Failed to delete grade", logger.Metadata{
			"error": err.Error(),
			"id":    id,
		})
		return err
	}
	return nil
}
