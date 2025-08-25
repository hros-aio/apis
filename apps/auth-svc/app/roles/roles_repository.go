package roles

import (
	"github.com/hros-aio/apis/libs/factory/function"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[RoleDB]
}

const REPOSITORY = "ROLE_REPOSITORY"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[RoleDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model *RoleModel) (*RoleModel, error) {
	input := model.DataMapper()
	if input.Code == "" {
		count, err := r.Model.Count(map[string]any{
			"tenant_id":  input.TenantID,
			"company_id": input.CompanyID,
		})
		if err != nil {
			return nil, err
		}
		input.Code, _ = function.GenerateCode("RO", count+1)
	}
	data, err := r.Model.Create(input)
	if err != nil {
		return nil, err
	}
	return data.Dto(), nil
}

func (r *Repository) FindAll(where sqlorm.Query, options sqlorm.FindOptions) ([]*RoleModel, int64, error) {
	data, total, err := r.Model.FindAllAndCount(where, options)
	if err != nil {
		return nil, 0, err
	}
	var result []*RoleModel
	for _, item := range data {
		result = append(result, item.Dto())
	}
	return result, total, nil
}

func (r *Repository) FindByID(id string) (*RoleModel, error) {
	data, err := r.Model.FindByID(id, sqlorm.FindOneOptions{
		Seperate: true,
		Related:  []string{"Company", "Permissions"},
	})
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, exception.NotFound("Permission not found")
	}
	return data.Dto(), nil
}

func (r *Repository) UpdateByID(id string, model *RoleModel) (*RoleModel, error) {
	data, err := r.Model.UpdateByID(id, model.DataMapper())
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, exception.NotFound("Permission not found")
	}
	return data.Dto(), nil
}
