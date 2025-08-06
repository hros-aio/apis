package tenant

import (
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[TenantDB]
}

const REPOSITORY = "TENANT_REPOSITORY"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[TenantDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model any) (*TenantModel, error) {
	data, err := r.Model.Create(model)
	if err != nil {
		return nil, err
	}
	return &TenantModel{
		Model:       base.Model{}.FromData(data.Model),
		Name:        data.Name,
		TenantId:    data.TenantId,
		Description: data.Description,
		Contact:     ContactPerson(data.Contact),
	}, nil
}

func (r *Repository) List() ([]TenantModel, error) {
	data, err := r.Model.FindAll(nil)
	if err != nil {
		return nil, err
	}
	tenants := []TenantModel{}
	for _, item := range data {
		tenants = append(tenants, TenantModel{
			Model:       base.Model{}.FromData(item.Model),
			Name:        item.Name,
			TenantId:    item.TenantId,
			Description: item.Description,
			Contact:     ContactPerson(item.Contact),
		})
	}

	return tenants, nil
}
