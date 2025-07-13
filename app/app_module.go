
package app

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewModule() core.Module {
	appModule := core.NewModule(core.NewModuleOptions{
		Controllers: []core.Controllers{NewController},
		Providers:   []core.Providers{NewService},
	})

	return appModule
}
