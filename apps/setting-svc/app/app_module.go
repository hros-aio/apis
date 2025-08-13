package app

import (
	"github.com/hros-aio/apis/apps/setting-svc/app/companies"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(&company.CompanyDB{}, &location.LocationDB{}),
			microservices.Register(),
			companies.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
		},
	})
}
