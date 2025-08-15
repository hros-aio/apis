package location

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[LocationDB]
}

const REPOSITORY = "LOCATION_REPOSITORY"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[LocationDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model *LocationModel) (*LocationModel, error) {
	input := model.DataMapper()
	data, err := r.Model.Create(input)
	if err != nil {
		return nil, err
	}
	return data.Dto(), nil
}

func (r *Repository) FindAll(where sqlorm.Query, options sqlorm.FindOptions) ([]*LocationModel, int64, error) {
	data, total, err := r.Model.FindAllAndCount(where, options)
	if err != nil {
		return nil, 0, err
	}
	models := []*LocationModel{}
	for _, val := range data {
		models = append(models, val.Dto())
	}

	return models, total, nil
}

func (r *Repository) FindByID(id string) (*LocationModel, error) {
	data, err := r.Model.FindByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, exception.NotFound("location not found")
	}
	return data.Dto(), nil
}

func (r *Repository) UpdateByID(id string, model *LocationModel) (*LocationModel, error) {
	input := model.DataMapper()
	data, err := r.Model.UpdateByID(id, input)
	if err != nil {
		return nil, err
	}

	return data.Dto(), nil
}
