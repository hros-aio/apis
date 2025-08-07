package user

import (
	"github.com/hros-aio/apis/libs/psql/common/base"
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
	return &UserModel{
		Model:      base.Model{}.FromData(data.Model),
		Username:   data.Username,
		TenantId:   data.TenantId,
		Password:   data.Password,
		Email:      data.Email,
		IsVerified: data.IsVerified,
		IsBanned:   data.IsBanned,
		IsAdmin:    data.IsAdmin,
	}, nil
}
