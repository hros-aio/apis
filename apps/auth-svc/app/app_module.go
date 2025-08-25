package app

import (
	"github.com/hros-aio/apis/apps/auth-svc/app/auth"
	"github.com/hros-aio/apis/apps/auth-svc/app/permissions"
	"github.com/hros-aio/apis/apps/auth-svc/app/shared"
	"github.com/hros-aio/apis/apps/auth-svc/app/users"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/psql/common/user"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(&user.UserDB{}, &location.LocationDB{}, &company.CompanyDB{}, &permissions.PermissionDB{}),
			saga.Register(),
			users.NewModule,
			auth.NewModule,
			shared.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
			middleware.AuthN,
		},
	})
}
