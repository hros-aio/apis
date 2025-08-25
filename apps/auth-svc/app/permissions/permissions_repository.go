package permissions

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[PermissionDB]
}

const REPOSITORY = "GRADE_REPOSITORY"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[PermissionDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model *PermissionModel) (*PermissionModel, error) {
	input := model.DataMapper()
	data, err := r.Model.Create(input)
	if err != nil {
		return nil, err
	}
	return data.Dto(), nil
}

func (r *Repository) FindAll(where sqlorm.Query, options sqlorm.FindOptions) ([]*PermissionModel, int64, error) {
	data, total, err := r.Model.FindAllAndCount(where, options)
	if err != nil {
		return nil, 0, err
	}
	var result []*PermissionModel
	for _, item := range data {
		result = append(result, item.Dto())
	}
	return result, total, nil
}

func (r *Repository) FindByID(id string) (*PermissionModel, error) {
	data, err := r.Model.FindByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, exception.NotFound("Permission not found")
	}
	return data.Dto(), nil
}

func (r *Repository) UpdateByID(id string, model *PermissionModel) (*PermissionModel, error) {
	data, err := r.Model.UpdateByID(id, model.DataMapper())
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, exception.NotFound("Permission not found")
	}
	return data.Dto(), nil
}
