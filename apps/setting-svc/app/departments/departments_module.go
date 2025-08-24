package departments

import (
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports:     []core.Modules{department.NewModule},
		Controllers: []core.Controllers{NewController},
		Providers:   []core.Providers{NewService},
	})
}
