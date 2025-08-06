package companies

import (
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("companies").
		Metadata(swagger.ApiTag("Companies")).
		Registry()

	ctrl.
		Pipe(core.BodyParser[CreateCompanyInput]{}).
		Post("", func(ctx core.Ctx) error {
			return ctx.JSON(nil)
		})

	return ctrl
}
