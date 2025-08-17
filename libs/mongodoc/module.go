package mongodoc

import (
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/config/v2"
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func Register(models ...any) core.Modules {
	return func(module core.Module) core.Module {
		return module.New(core.NewModuleOptions{
			Imports: []core.Modules{
				mongoose.ForRootFactory(func(ref core.Module) *mongoose.Connect {
					cfg := config.Inject[shared.Config](ref)

					return mongoose.New(cfg.Mongo.Url)
				}),
			},
		})
	}
}
