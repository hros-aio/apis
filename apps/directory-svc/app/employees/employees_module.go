package employees

import (
	"github.com/hros-aio/apis/libs/psql/common/employee"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports:     []core.Modules{employee.NewModule},
		Controllers: []core.Controllers{NewController},
	})
}
