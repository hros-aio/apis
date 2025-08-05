package company

import (
	"github.com/hros-aio/apis/libs/sql/common/base"
	"github.com/hros-aio/apis/libs/sql/common/tenant"
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

func (r *Repository) Create(model any) (*CompanyModel, error) {
	data, err := r.Model.Create(model)
	if err != nil {
		return nil, err
	}

	return &CompanyModel{
		Model:            base.Model{}.FromData(data.Model),
		TenantId:         data.TenantId,
		Name:             data.Name,
		Industry:         data.Industry,
		Size:             data.Size,
		Logo:             data.Logo,
		Contact:          tenant.ContactPerson(data.Contact),
		SecondaryContact: tenant.ContactPerson(data.SecondaryContact),
	}, nil
}
