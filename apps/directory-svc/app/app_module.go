package app

import (
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/psql"
	"github.com/hros-aio/apis/libs/psql/common/employee"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			psql.Register(&employee.EmployeeDB{}),
			saga.Register(),
		},
	})
}
