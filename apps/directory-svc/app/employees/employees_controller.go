package employees

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	svc := core.Inject[EmployeeService](module)
	ctrl := module.NewController("employees").
		Metadata(swagger.ApiTag("Employee"), swagger.ApiSecurity("bearerAuth")).
		Registry()

	ctrl.Pipe(core.BodyParser[CreateEmployeeInput]{}).Post("", func(ctx core.Ctx) error {
		contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
		input := core.Execution[CreateEmployeeInput](core.InBody, ctx)

		model := input.Dto(*contextInfo)
		data, err := svc.Create(*contextInfo, model)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	})

	return ctrl
}
