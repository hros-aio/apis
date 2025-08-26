package locations

import (
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports:   []core.Modules{location.NewModule},
		Providers: []core.Providers{NewHandler},
	})
}
