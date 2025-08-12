package companies

import (
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type CompanyService struct {
	companyRepo *company.Repository
}

func NewService(module core.Module) core.Provider {
	companyRepo := module.Ref(company.REPOSITORY).(*company.Repository)

	return module.NewProvider(&CompanyService{
		companyRepo: companyRepo,
	})
}
