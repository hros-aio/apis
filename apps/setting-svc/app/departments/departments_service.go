package departments

import (
	"reflect"

	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type DepartmentService struct {
	departmentRepo *department.Repository
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
}

func NewService(module core.Module) core.Provider {
	departmentRepo := module.Ref(department.REPOSITORY).(*department.Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&DepartmentService{
		departmentRepo: departmentRepo,
		logger:         logger,
		eventPublisher: eventPublisher,
	})
}

func (s *DepartmentService) Create(ctx middleware.ContextInfo, model *department.DepartmentModel) (*department.DepartmentModel, error) {
	if reflect.ValueOf(model.TenantID).IsZero() {
		model.TenantID = ctx.TenantID
	}
	if reflect.ValueOf(model.CompanyID).IsZero() {
		model.CompanyID = ctx.CompanyID
	}
	createdLocation, err := s.departmentRepo.Create(model)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return createdLocation, nil
}

func (s *DepartmentService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*department.DepartmentModel, int64, error) {
	filter := make(map[string]any)

	if ctx.TenantID != "" {
		filter["tenant_id"] = ctx.TenantID
	}

	if !reflect.ValueOf(ctx.CompanyID).IsZero() {
		filter["company_id"] = ctx.CompanyID
	}
	data, total, err := s.departmentRepo.FindAll(filter, sqlorm.FindOptions{
		Offset: queryParams.Skip,
		Limit:  queryParams.Limit,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	return data, total, nil
}

func (s *DepartmentService) GetByID(ctx middleware.ContextInfo, id string) (*department.DepartmentModel, error) {
	foundLocation, err := s.departmentRepo.FindByID(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return foundLocation, nil
}

func (s *DepartmentService) UpdateByID(ctx middleware.ContextInfo, id string, model *department.DepartmentModel) (*department.DepartmentModel, error) {
	updatedLocation, err := s.departmentRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return updatedLocation, nil
}

func (s *DepartmentService) DeleteById(ctx middleware.ContextInfo, id string) error {
	err := s.departmentRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
