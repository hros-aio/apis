package app

import (
	"github.com/hros-aio/apis/apps/time-svc/app/shared"
	"github.com/hros-aio/apis/apps/time-svc/app/work_schedules"
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/mongodoc"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/tinh-tinh/pubsub/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	appModule := core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			mongodoc.Register(),
			saga.Register(),
			microservices.Register(),
			pubsub.ForRoot(pubsub.BrokerOptions{
				MaxSubscribers: 100,
			}),
			shared.NewModule,
			work_schedules.NewModule,
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
			middleware.AuthN,
		},
	})

	return appModule
}
