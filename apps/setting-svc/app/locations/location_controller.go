package locations

import (
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("locations").
		Metadata(swagger.ApiTag("Location"), swagger.ApiSecurity("bearerAuth")).
		Registry()

	ctrl.Pipe(
		core.QueryParser[shared.QueryCompany]{},
		core.BodyParser[CreateLocationInput]{},
		core.PathParser[shared.ParamID]{},
	).
		Post("", func(ctx core.Ctx) error {
			return ctx.JSON(nil)
		})

	return ctrl
}
