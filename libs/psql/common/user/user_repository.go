package user

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Repository struct {
	Model *sqlorm.Repository[UserDB]
}

const REPOSITORY = "USER_REPOSITORY	"

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[UserDB](module)
	return module.NewProvider(core.ProviderOptions{
		Name: REPOSITORY,
		Value: &Repository{
			Model: repo,
		},
	})
}

func (r *Repository) Create(model any) (*UserModel, error) {
	data, err := r.Model.Create(model)
	if err != nil {
		return nil, err
	}
	return data.Dto(), nil
}

func (r *Repository) FindByEmail(email string) (*UserModel, error) {
	data, err := r.Model.FindOne(map[string]any{
		"email": email,
	})
	if err != nil {
		return nil, err
	}

	return data.Dto(), nil
}

func (r *Repository) FindByID(id string) (*UserModel, error) {
	data, err := r.Model.FindByID(id)
	if err != nil {
		return nil, err
	}
	data.Password = ""
	return data.Dto(), nil
}
