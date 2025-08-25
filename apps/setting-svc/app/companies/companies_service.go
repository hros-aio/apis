package companies

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type CompanyService struct {
	companyRepo    *company.Repository
	logger         *logger.Logger
	eventPublisher *saga.EventPulisher
}

func NewService(module core.Module) core.Provider {
	companyRepo := module.Ref(company.REPOSITORY).(*company.Repository)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	return module.NewProvider(&CompanyService{
		companyRepo:    companyRepo,
		logger:         logger,
		eventPublisher: eventPublisher,
	})
}

func (s *CompanyService) Create(ctx middleware.ContextInfo, model *company.CompanyModel) (*company.CompanyModel, error) {
	if model.TenantID == "" {
		model.TenantID = ctx.TenantID
	}
	createdCompany, err := s.companyRepo.Create(model)
	if err != nil {
		s.logger.Error("Failed to create company", logger.Metadata{
			"error": err.Error(),
			"model": model,
		})
		return nil, err
	}
	return createdCompany, nil
}

func (s *CompanyService) List(ctx middleware.ContextInfo, queryParams middleware.Paginate) ([]*company.CompanyModel, int64, error) {
	data, total, err := s.companyRepo.FindAll(map[string]any{
		"tenant_id": ctx.TenantID,
	}, sqlorm.FindOptions{
		Offset: queryParams.Skip,
		Limit:  queryParams.Limit,
	})
	if err != nil {
		s.logger.Error("Failed to fetch list companies", logger.Metadata{
			"error": err.Error(),
			"query": queryParams,
		})
		return nil, 0, err
	}

	return data, total, nil
}

func (s *CompanyService) UpdateById(ctx middleware.ContextInfo, id string, model *company.CompanyModel) (*company.CompanyModel, error) {
	foundCompany, err := s.companyRepo.FindByID(id)
	if err != nil {
		s.logger.Error("Failed to find company", logger.Metadata{
			"error": err.Error(),
		})
		return nil, err
	}
	if err := s.checkTenantID(ctx, foundCompany); err != nil {
		s.logger.Error("Not allowed access", logger.Metadata{
			"tenant_id": ctx.TenantID,
		})
		return nil, err
	}

	updatedCompany, err := s.companyRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to update company", logger.Metadata{
			"error": err.Error(),
			"model": model,
		})
		return nil, err
	}

	return updatedCompany, nil
}

func (s *CompanyService) ActiveByID(ctx middleware.ContextInfo, id string) (*company.CompanyModel, error) {
	foundCompany, err := s.companyRepo.FindByID(id)
	if err != nil {
		s.logger.Error("Failed to find company", logger.Metadata{
			"error": err.Error(),
		})
		return nil, err
	}

	if err := s.checkTenantID(ctx, foundCompany); err != nil {
		s.logger.Error("Not allowed access", logger.Metadata{
			"tenant_id": ctx.TenantID,
		})
		return nil, err
	}

	model := &company.CompanyModel{Status: company.ActiveStatus}
	activeCompany, err := s.companyRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to activate company", logger.Metadata{
			"error": err.Error(),
			"model": model,
		})
		return nil, err
	}

	go s.eventPublisher.RegisterSync(events.CompanyActivated, messages.SyncDataPayload{
		Data: ToActiveMessage(activeCompany),
	})
	return activeCompany, nil
}

func (s *CompanyService) DeactiveByID(ctx middleware.ContextInfo, id string) (*company.CompanyModel, error) {
	foundCompany, err := s.companyRepo.FindByID(id)
	if err != nil {
		s.logger.Error("Failed to find company", logger.Metadata{
			"error": err.Error(),
		})
		return nil, err
	}

	if err := s.checkTenantID(ctx, foundCompany); err != nil {
		s.logger.Error("Not allowed access", logger.Metadata{
			"tenant_id": ctx.TenantID,
		})
		return nil, err
	}

	model := &company.CompanyModel{Status: company.InactiveStatus}
	deactiveCompany, err := s.companyRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to deactivate company", logger.Metadata{
			"error": err.Error(),
			"model": model,
		})
		return nil, err
	}

	return deactiveCompany, nil
}

func (s *CompanyService) checkTenantID(ctx middleware.ContextInfo, data *company.CompanyModel) error {
	if data.TenantID != ctx.TenantID {
		return exception.Forbidden("You do not have permission to access this resource")
	}
	return nil
}
