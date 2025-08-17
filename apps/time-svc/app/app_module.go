package app

import (
	"github.com/hros-aio/apis/libs/factory"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	appModule := core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
		},
	})

	return appModule
}
