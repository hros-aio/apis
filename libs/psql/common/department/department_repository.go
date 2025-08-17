package department

import (
	"github.com/hros-aio/apis/libs/factory/function"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[DepartmentDB]
}

const REPOSITORY = "DEPARTMENT_REPOSITORY"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[DepartmentDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model *DepartmentModel) (*DepartmentModel, error) {
	input := model.DataMapper()
	if input.Code == "" {
		count, err := r.Model.Count(map[string]any{
			"tenant_id":  input.TenantID,
			"company_id": input.CompanyID,
		})
		if err != nil {
			return nil, err
		}
		input.Code, _ = function.GenerateCode("DP", count+1)
	}
	data, err := r.Model.Create(input)
	if err != nil {
		return nil, err
	}
	return data.Dto(), nil
}

func (r *Repository) FindAll(where sqlorm.Query, options sqlorm.FindOptions) ([]*DepartmentModel, int64, error) {
	data, total, err := r.Model.FindAllAndCount(where, options)
	if err != nil {
		return nil, 0, err
	}
	models := []*DepartmentModel{}
	for _, val := range data {
		models = append(models, val.Dto())
	}

	return models, total, nil
}

func (r *Repository) FindByID(id string) (*DepartmentModel, error) {
	data, err := r.Model.FindByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, exception.NotFound("location not found")
	}
	return data.Dto(), nil
}

func (r *Repository) UpdateByID(id string, model *DepartmentModel) (*DepartmentModel, error) {
	input := model.DataMapper()
	data, err := r.Model.UpdateByID(id, input)
	if err != nil {
		return nil, err
	}

	return data.Dto(), nil
}
