package location

import (
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports: []core.Modules{
			mongoose.ForFeature(mongoose.NewModel[LocationSchema]()),
		},
		Providers: []core.Providers{NewRepository},
		Exports:   []core.Providers{NewRepository},
	})
}
