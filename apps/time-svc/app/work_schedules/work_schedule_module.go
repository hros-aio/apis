package works_chedules

import (
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports: []core.Modules{
			mongoose.ForFeature(mongoose.NewModel[WorkScheduleSchema]()),
		},
		Controllers: []core.Controllers{NewController},
		Providers:   []core.Providers{NewRepository, NewService, InternalHandler},
		Exports:     []core.Providers{NewRepository},
	})
}
