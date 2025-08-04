package company

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[CompanyDB]
}

const REPOSITORY = "COMPANY_REPOSITORY"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[CompanyDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model any) {}
