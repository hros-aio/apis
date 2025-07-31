package tenant

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type TenantRepository struct {
	repo *sqlorm.Repository[TenantModel]
}

func NewRepository(module core.Module) core.Provider {
	repo := sqlorm.InjectRepository[TenantModel](module)
	return module.NewProvider(&TenantRepository{
		repo: repo,
	})
}
