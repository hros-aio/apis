package companies

import (
	"github.com/hros-aio/apis/libs/sql/common/company"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports:     []core.Modules{company.NewModule},
		Controllers: []core.Controllers{NewController},
		Providers:   []core.Providers{NewService},
	})
}
