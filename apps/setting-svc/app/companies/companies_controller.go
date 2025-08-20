package companies

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	svc := core.Inject[CompanyService](module)
	ctrl := module.NewController("companies").
		Metadata(swagger.ApiTag("Company"), swagger.ApiSecurity("bearerAuth")).
		Registry()

	ctrl.
		Pipe(core.BodyParser[CreateCompanyInput]{}).
		Post("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[CreateCompanyInput](core.InBody, ctx)

			model := input.Dto()
			data, err := svc.Create(*contextInfo, model)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	ctrl.
		Use(middleware.Pagination).
		Get("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			queryParams := core.Execution[middleware.Paginate](middleware.PAGINATE, ctx)
			data, total, err := svc.List(*contextInfo, *queryParams)
			if err != nil {
				return err
			}
			return ctx.JSON(core.Map{
				"data":  data,
				"total": total,
			})
		})

	ctrl.Pipe(
		core.BodyParser[UpdateCompanyInput]{},
		core.PathParser[shared.ParamID]{},
	).Put("{id}", func(ctx core.Ctx) error {
		contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
		input := core.Execution[UpdateCompanyInput](core.InBody, ctx)

		model := input.Dto()
		data, err := svc.UpdateById(*contextInfo, ctx.Path("id"), model)
		if err != nil {
			return err
		}
		return ctx.JSON(core.Map{
			"data": data,
		})
	})

	ctrl.
		Pipe(core.PathParser[shared.ParamID]{}).
		Delete("{id}", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)

			err := svc.DeleteById(*contextInfo, ctx.Path("id"))
			if err != nil {
				return err
			}
			return ctx.JSON(core.Map{
				"status": true,
			})
		})

	return ctrl
}
