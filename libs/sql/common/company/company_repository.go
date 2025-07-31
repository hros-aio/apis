package company

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type CompanyRepository struct {
	repo *sqlorm.Repository[CompanyModel]
}

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[CompanyModel](module)
	return module.NewProvider(&CompanyRepository{
		repo: repo,
	})
}
