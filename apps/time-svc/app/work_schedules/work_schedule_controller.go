package works_chedules

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	svc := core.Inject[WorkScheduleService](module)
	ctrl := module.NewController("work-schedules").
		Metadata(swagger.ApiTag("Work Schedule"), swagger.ApiSecurity("bearerAuth")).
		Registry()

	ctrl.
		Pipe(core.BodyParser[CreateWorkScheduleInput]{}).
		Post("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[CreateWorkScheduleInput](core.InBody, ctx)

			model := input.Dto()
			data, err := svc.Create(*contextInfo, model)
			if err != nil {
				return err
			}

			return ctx.JSON(data)
		})

	ctrl.
		Pipe(core.QueryParser[middleware.PaginateInput]{}).
		Use(middleware.Pagination).
		Get("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			queryParams := core.Execution[middleware.Paginate](middleware.PAGINATE, ctx)

			data, err := svc.List(*contextInfo, *queryParams)
			if err != nil {
				return err
			}

			return ctx.JSON(data)
		})

	ctrl.Put("/:id", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	ctrl.Delete("/:id", func(ctx core.Ctx) error {
		return ctx.JSON(nil)
	})

	return ctrl
}
