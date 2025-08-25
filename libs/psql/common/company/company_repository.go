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

func (r *Repository) Create(model *CompanyModel) (*CompanyModel, error) {
	input := model.DataMapper()
	data, err := r.Model.Create(input)
	if err != nil {
		return nil, err
	}

	return data.Dto(), nil
}

func (r *Repository) FindAll(where sqlorm.Query, options sqlorm.FindOptions) ([]*CompanyModel, int64, error) {
	data, total, err := r.Model.FindAllAndCount(where, options)
	if err != nil {
		return nil, 0, err
	}
	models := []*CompanyModel{}
	for _, val := range data {
		models = append(models, val.Dto())
	}

	return models, total, nil
}

func (r *Repository) FindByID(id string) (*CompanyModel, error) {
	data, err := r.Model.FindByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	return data.Dto(), nil
}

func (r *Repository) UpdateByID(id string, model *CompanyModel) (*CompanyModel, error) {
	input := model.DataMapper()
	_, err := r.Model.UpdateByID(id, input)
	if err != nil {
		return nil, err
	}

	data, err := r.Model.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data.Dto(), nil
}
