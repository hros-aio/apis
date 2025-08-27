package departments

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
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
	createdDepartment, err := s.departmentRepo.Create(model)
	if err != nil {
		s.logger.Error("[DepartmentService][Create] Failed to create department", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}

	go s.eventPublisher.RegisterSync(events.DepartmentCreated, ToCreatedMessage(createdDepartment))
	return createdDepartment, nil
}

func (s *DepartmentService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*department.DepartmentModel, int64, error) {
	filter := make(map[string]any)

	if ctx.TenantID != "" {
		filter["tenant_id"] = ctx.TenantID
	}

	if ctx.CompanyID != uuid.Nil {
		filter["company_id"] = ctx.CompanyID
	}
	data, total, err := s.departmentRepo.FindAll(filter, sqlorm.FindOptions{
		Offset: queryParams.Skip,
		Limit:  queryParams.Limit,
	})
	if err != nil {
		s.logger.Error("[DepartmentService][List] Failed to list departments", logger.Metadata{
			"err": err.Error(),
		})
		return nil, 0, err
	}

	return data, total, nil
}

func (s *DepartmentService) GetByID(ctx middleware.ContextInfo, id string) (*department.DepartmentModel, error) {
	foundDepartment, err := s.departmentRepo.FindByID(id)
	if err != nil {
		s.logger.Error("[DepartmentService][GetByID] Failed to get department by ID", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}

	return foundDepartment, nil
}

func (s *DepartmentService) UpdateByID(ctx middleware.ContextInfo, id string, model *department.DepartmentModel) (*department.DepartmentModel, error) {
	foundDepartment, err := s.departmentRepo.FindByID(id)
	if err != nil {
		s.logger.Error("[DepartmentService][UpdateByID] Failed to find department by ID", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}
	if foundDepartment == nil {
		s.logger.Error("[DepartmentService][UpdateByID] Department not found", logger.Metadata{
			"id": id,
		})
		return nil, exception.NotFound("department not found")
	}

	updatedDepartment, err := s.departmentRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("[DepartmentService][UpdateByID] Failed to update department by ID", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}

	go s.eventPublisher.RegisterSync(events.DepartmentUpdated, ToUpdatedMessage(foundDepartment, updatedDepartment))
	return updatedDepartment, nil
}

func (s *DepartmentService) DeleteById(ctx middleware.ContextInfo, id string) error {
	err := s.departmentRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error("[DepartmentService][DeleteById] Failed to delete department by ID", logger.Metadata{
			"err": err.Error(),
		})
		return err
	}
	return nil
}
