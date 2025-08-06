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
