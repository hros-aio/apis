package app

import (
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/workers/sync-worker/app/sync"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			saga.Register(),
			microservices.Register(),
			sync.NewModule,
		},
	})
}
