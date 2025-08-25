package shared

import (
	"github.com/hros-aio/apis/apps/auth-svc/app/shared/companies"
	"github.com/hros-aio/apis/apps/auth-svc/app/shared/locations"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule(module core.Module) core.Module {
	return module.New(core.NewModuleOptions{
		Imports: []core.Modules{
			companies.NewModule,
			locations.NewModule,
		},
	})
}
