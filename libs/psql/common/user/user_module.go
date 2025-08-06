package user

import (
	"github.com/tinh-tinh/sqlorm/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports: []core.Modules{
			sqlorm.ForFeature(sqlorm.NewRepo(UserDB{})),
		},
		Providers: []core.Providers{NewRepository},
		Exports:   []core.Providers{NewRepository},
	})
}
