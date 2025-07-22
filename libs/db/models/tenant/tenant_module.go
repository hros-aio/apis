package tenant

import (
	"github.com/hros-aio/apis/libs/db/models/tenant/command"
	"github.com/hros-aio/apis/libs/db/models/tenant/query"
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports: []core.Modules{
			sqlorm.ForFeature(sqlorm.NewRepo(command.TenantDB{})),
			mongoose.ForFeature(mongoose.NewModel[query.TenantSchema]("tenants")),
		},
		Providers: []core.Providers{command.NewRepository, query.NewHandler},
		Exports:   []core.Providers{command.NewRepository, query.NewHandler},
	})
}
