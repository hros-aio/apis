package tenants

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
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
			contextInfo, ok := ctx.Get(middleware.APP_CONTEXT).(middleware.ContextInfo)
			if !ok {
				return exception.InternalServer("Cannot parse context")
			}
			input, ok := ctx.Body().(*TenantCreateInput)
			if !ok {
				return exception.BadRequest("Cannot parse input")
			}
			data, err := svc.Create(contextInfo, input)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	ctrl.Get("", func(ctx core.Ctx) error {
		contextInfo, ok := ctx.Get(middleware.APP_CONTEXT).(middleware.ContextInfo)
		if !ok {
			return exception.BadRequest("Cannot parse input")
		}
		data, err := svc.List(contextInfo)
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
