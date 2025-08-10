package app

import (
	"github.com/hros-aio/apis/apps/auth-svc/app/auth"
	"github.com/hros-aio/apis/apps/auth-svc/app/users"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/user"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(&user.UserDB{}),
			microservices.Register(),
			users.NewModule,
			auth.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
			middleware.AuthN,
		},
	})
}
