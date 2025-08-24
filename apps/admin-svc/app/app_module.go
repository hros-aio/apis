package app

import (
	"github.com/hros-aio/apis/apps/admin-svc/app/tenants"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(&tenant.TenantDB{}),
			saga.Register(),
			tenants.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
		},
	})
}
