package tenant

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports:     []core.Modules{},
		Controllers: []core.Controllers{NewController},
	})
}
