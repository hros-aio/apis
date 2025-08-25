package permissions

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type PermissionService struct {
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
	permissionRepo *Repository
}

func NewService(module core.Module) core.Provider {
	permissionRepo := module.Ref(REPOSITORY).(*Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&PermissionService{
		eventPublisher: eventPublisher,
		logger:         logger,
		permissionRepo: permissionRepo,
	})
}

func (s *PermissionService) Create(ctx middleware.ContextInfo, model *PermissionModel) (*PermissionModel, error) {
	created, err := s.permissionRepo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create permission", logger.Metadata{
			"err":   err.Error(),
			"input": model,
		})
		return nil, err
	}
	return created, nil
}

func (s *PermissionService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*PermissionModel, int64, error) {
	data, total, err := s.permissionRepo.FindAll(nil, sqlorm.FindOptions{
		Limit:  queryParams.Limit,
		Offset: queryParams.Skip,
	})
	if err != nil {
		s.logger.Error("Failed to fetch list permissions", logger.Metadata{
			"error": err.Error(),
		})
		return nil, 0, err
	}
	return data, total, nil
}
func (s *PermissionService) UpdateByID(ctx middleware.ContextInfo, id string, model *PermissionModel) (*PermissionModel, error) {
	updatedPermission, err := s.permissionRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to update permission", logger.Metadata{
			"error": err.Error(),
			"input": model,
		})
		return nil, err
	}
	return updatedPermission, nil
}

func (s *PermissionService) DeleteByID(ctx middleware.ContextInfo, id string) error {
	err := s.permissionRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error("Failed to delete permission", logger.Metadata{
			"error": err.Error(),
			"id":    id,
		})
		return err
	}
	return nil
}
