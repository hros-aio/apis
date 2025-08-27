package locations

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type LocationService struct {
	locationRepo   *location.Repository
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
}

func NewService(module core.Module) core.Provider {
	locationRepo := module.Ref(location.REPOSITORY).(*location.Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&LocationService{
		locationRepo:   locationRepo,
		logger:         logger,
		eventPublisher: eventPublisher,
	})
}

func (s *LocationService) Create(ctx middleware.ContextInfo, model *location.LocationModel) (*location.LocationModel, error) {
	createdLocation, err := s.locationRepo.Create(model)
	if err != nil {
		s.logger.Error("[LocationService][Create] Failed to create location", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}

	go s.eventPublisher.RegisterSync(events.LocationCreated, ToCreatedMessage(createdLocation))
	return createdLocation, nil
}

func (s *LocationService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*location.LocationModel, int64, error) {
	filter := make(map[string]any)

	if ctx.TenantID != "" {
		filter["tenant_id"] = ctx.TenantID
	}

	if ctx.CompanyID != uuid.Nil {
		filter["company_id"] = ctx.CompanyID
	}
	data, total, err := s.locationRepo.FindAll(filter, sqlorm.FindOptions{
		Offset: queryParams.Skip,
		Limit:  queryParams.Limit,
	})
	if err != nil {
		s.logger.Error("[LocationService][List] Failed to list locations", logger.Metadata{
			"err": err.Error(),
		})
		return nil, 0, err
	}

	return data, total, nil
}

func (s *LocationService) GetByID(ctx middleware.ContextInfo, id string) (*location.LocationModel, error) {
	foundLocation, err := s.locationRepo.FindByID(id)
	if err != nil {
		s.logger.Error("[LocationService][GetByID] Failed to get location by ID", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}

	return foundLocation, nil
}

func (s *LocationService) UpdateByID(ctx middleware.ContextInfo, id string, model *location.LocationModel) (*location.LocationModel, error) {
	foundLocation, err := s.locationRepo.FindByID(id)
	if err != nil {
		s.logger.Error("[LocationService][UpdateByID] Failed to find location by ID", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}
	if foundLocation == nil {
		s.logger.Error("[LocationService][UpdateByID] Location not found", logger.Metadata{
			"id": id,
		})
		return nil, exception.NotFound("location not found")
	}

	updatedLocation, err := s.locationRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("[LocationService][UpdateByID] Failed to update location by ID", logger.Metadata{
			"err": err.Error(),
		})
		return nil, err
	}
	go s.eventPublisher.Publish(events.LocationUpdated, ToUpdatedMessage(foundLocation, updatedLocation))
	return updatedLocation, nil
}

func (s *LocationService) DeleteById(ctx middleware.ContextInfo, id string) error {
	err := s.locationRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	go s.eventPublisher.Publish(events.LocationDeleted, messages.LocationDeletedPayload{
		Id:        id,
		TenantID:  ctx.TenantID,
		CompanyID: ctx.CompanyID.String(),
	})
	return nil
}
