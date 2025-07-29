package tenant

import (
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("tenants").
		Metadata(swagger.ApiTag("Tenants")).
		Registry()

	ctrl.Post("", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	ctrl.Get("", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
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
