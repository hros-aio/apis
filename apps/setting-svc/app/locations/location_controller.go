package locations

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	svc := core.Inject[LocationService](module)
	ctrl := module.NewController("locations").
		Metadata(swagger.ApiTag("Location"), swagger.ApiSecurity("bearerAuth")).
		Registry()

	ctrl.Pipe(
		core.QueryParser[shared.QueryCompany]{},
		core.BodyParser[CreateLocationInput]{},
		core.PathParser[shared.ParamID]{},
	).
		Post("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[CreateLocationInput](core.InBody, ctx)

			model := input.Dto()
			data, err := svc.Create(*contextInfo, model)
			if err != nil {
				return err
			}
			return ctx.JSON(core.Map{
				"data": data,
			})
		})

	return ctrl
}
