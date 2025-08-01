package app

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{},
	})
}
