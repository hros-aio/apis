package app

import (
	"github.com/hros-aio/apis/apps/directory-svc/app/employees"
	"github.com/hros-aio/apis/apps/directory-svc/app/shared"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/employee"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	models := []interface{}{
		&company.CompanyDB{},
		&location.LocationDB{},
		&employee.EmployeeDB{},
	}
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(models...),
			saga.Register(),
			shared.NewModule,
			employees.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
			middleware.AuthN,
		},
	})
}
