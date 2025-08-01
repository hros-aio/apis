package app

import (
	"github.com/hros-aio/apis/apps/admin-svc/app/tenant"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			tenant.NewModule,
		},
	})
}
