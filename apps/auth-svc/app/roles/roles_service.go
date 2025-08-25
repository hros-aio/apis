package roles

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type RoleService struct {
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
	roleRepo       *Repository
}

func NewService(module core.Module) core.Provider {
	roleRepo := module.Ref(REPOSITORY).(*Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&RoleService{
		eventPublisher: eventPublisher,
		logger:         logger,
		roleRepo:       roleRepo,
	})
}

func (s *RoleService) Create(ctx middleware.ContextInfo, model *RoleModel) (*RoleModel, error) {
	created, err := s.roleRepo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create role", logger.Metadata{
			"err":   err.Error(),
			"input": model,
		})
		return nil, err
	}
	return created, nil
}

func (s *RoleService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*RoleModel, int64, error) {
	data, total, err := s.roleRepo.FindAll(nil, sqlorm.FindOptions{
		Limit:  queryParams.Limit,
		Offset: queryParams.Skip,
	})
	if err != nil {
		s.logger.Error("Failed to fetch list roles", logger.Metadata{
			"error": err.Error(),
		})
		return nil, 0, err
	}
	return data, total, nil
}

func (s *RoleService) GetByID(ctx middleware.ContextInfo, id string) (*RoleModel, error) {
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		s.logger.Error("Failed to fetch role by ID", logger.Metadata{
			"error": err.Error(),
			"id":    id,
		})
		return nil, err
	}
	return role, nil
}

func (s *RoleService) UpdateByID(ctx middleware.ContextInfo, id string, model *RoleModel) (*RoleModel, error) {
	updatedRole, err := s.roleRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to update role", logger.Metadata{
			"error": err.Error(),
			"input": model,
		})
		return nil, err
	}
	return updatedRole, nil
}

func (s *RoleService) DeleteByID(ctx middleware.ContextInfo, id string) error {
	err := s.roleRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error("Failed to delete role", logger.Metadata{
			"error": err.Error(),
			"id":    id,
		})
		return err
	}
	return nil
}
