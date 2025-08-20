package app

import (
	"github.com/hros-aio/apis/apps/setting-svc/app/companies"
	"github.com/hros-aio/apis/apps/setting-svc/app/departments"
	"github.com/hros-aio/apis/apps/setting-svc/app/grades"
	"github.com/hros-aio/apis/apps/setting-svc/app/locations"
	"github.com/hros-aio/apis/apps/setting-svc/app/titles"
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
			psql.Register(&company.CompanyDB{}, &location.LocationDB{}, &department.DepartmentDB{}, &grades.GradeDB{}, &titles.TitleDB{}),
			saga.Register(),
			companies.NewModule,
			locations.NewModule,
			departments.NewModule,
			grades.NewModule,
			titles.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
			middleware.AuthN,
		},
	})
}
