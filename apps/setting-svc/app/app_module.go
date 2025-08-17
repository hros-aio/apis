package app

import (
	"github.com/hros-aio/apis/apps/setting-svc/app/companies"
	"github.com/hros-aio/apis/apps/setting-svc/app/departments"
	"github.com/hros-aio/apis/apps/setting-svc/app/locations"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(&company.CompanyDB{}, &location.LocationDB{}, &department.DepartmentDB{}),
			saga.Register(),
			companies.NewModule,
			locations.NewModule,
			departments.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
			middleware.AuthN,
		},
	})
}
