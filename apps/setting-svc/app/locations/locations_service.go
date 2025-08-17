package locations

import (
	"reflect"

	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/sqlorm/v2"
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
	if reflect.ValueOf(model.TenantId).IsZero() {
		model.TenantId = ctx.TenantID
	}
	if reflect.ValueOf(model.CompanyID).IsZero() {
		model.CompanyID = ctx.CompanyID
	}
	createdLocation, err := s.locationRepo.Create(model)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	go s.eventPublisher.Publish(events.LocationCreated, ToCreatedMessage(createdLocation))
	return createdLocation, nil
}

func (s *LocationService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*location.LocationModel, int64, error) {
	filter := make(map[string]any)

	if ctx.TenantID != "" {
		filter["tenant_id"] = ctx.TenantID
	}

	if !reflect.ValueOf(ctx.CompanyID).IsZero() {
		filter["company_id"] = ctx.CompanyID
	}
	data, total, err := s.locationRepo.FindAll(filter, sqlorm.FindOptions{
		Offset: queryParams.Skip,
		Limit:  queryParams.Limit,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	return data, total, nil
}

func (s *LocationService) GetByID(ctx middleware.ContextInfo, id string) (*location.LocationModel, error) {
	foundLocation, err := s.locationRepo.FindByID(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return foundLocation, nil
}

func (s *LocationService) UpdateByID(ctx middleware.ContextInfo, id string, model *location.LocationModel) (*location.LocationModel, error) {
	foundLocation, err := s.locationRepo.FindByID(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	updatedLocation, err := s.locationRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error(err.Error())
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
