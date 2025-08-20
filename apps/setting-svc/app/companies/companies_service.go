package companies

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type CompanyService struct {
	companyRepo *company.Repository
	logger      *logger.Logger
}

func NewService(module core.Module) core.Provider {
	companyRepo := module.Ref(company.REPOSITORY).(*company.Repository)
	logger := logger.InjectLog(module)

	return module.NewProvider(&CompanyService{
		companyRepo: companyRepo,
		logger:      logger,
	})
}

func (s *CompanyService) Create(ctx middleware.ContextInfo, model *company.CompanyModel) (*company.CompanyModel, error) {
	if model.TenantId == "" {
		model.TenantId = ctx.TenantID
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
	foundCompany, err := s.companyRepo.UpdateByID(id, model)
	if err != nil {
		s.logger.Error("Failed to update company", logger.Metadata{
			"error": err.Error(),
			"model": model,
		})
		return nil, err
	}
	return foundCompany, nil
}

func (s *CompanyService) DeleteById(ctx middleware.ContextInfo, id string) error {
	err := s.companyRepo.Model.DeleteByID(id)
	if err != nil {
		s.logger.Error("Failed to delete company", logger.Metadata{
			"error": err.Error(),
		})
		return err
	}
	return nil
}
