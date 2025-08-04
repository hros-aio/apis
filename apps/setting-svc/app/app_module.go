package app

import (
	"github.com/hros-aio/apis/libs/factory"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/sql"
	"github.com/hros-aio/apis/libs/sql/common/company"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewModule() core.Module {
	return core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			factory.Register(),
			sql.Register(&company.CompanyDB{}),
		},
		Middlewares: []core.Middleware{
			middleware.SetContext,
		},
	})
}
