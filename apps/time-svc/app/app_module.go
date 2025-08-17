package app

import (
	"time-svc/app/shared"

	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/mongodoc"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	appModule := core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			mongodoc.Register(),
			microservices.Register(),
			shared.NewModule,
		},
	})

	return appModule
}
