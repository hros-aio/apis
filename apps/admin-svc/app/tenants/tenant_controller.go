package tenants

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("tenants").
		Metadata(swagger.ApiTag("Tenants")).
		Registry()

	svc := core.Inject[TenantService](module)
	ctrl.
		Pipe(core.BodyParser[TenantCreateInput]{}).
		Post("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[TenantCreateInput](core.InBody, ctx)

			data, err := svc.Create(*contextInfo, input)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	ctrl.Get("", func(ctx core.Ctx) error {
		contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
		data, err := svc.List(*contextInfo)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	})

	ctrl.Get("sponsors", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	ctrl.Get("{id}", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	ctrl.Put("{id}", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	ctrl.Delete("{id}", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	return ctrl
}
