package app

import (
	"github.com/hros-aio/apis/libs/db"
	"github.com/hros-aio/apis/libs/db/models/tenant"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			db.Register(),
			tenant.NewModule,
		},
	})
}
