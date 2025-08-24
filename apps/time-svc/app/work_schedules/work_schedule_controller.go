package work_schedules

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/factory/shared"
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

			model := input.Dto(*contextInfo)
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

	ctrl.
		Pipe(
			core.PathParser[shared.ParamID]{},
			core.BodyParser[UpdateWorkScheduleInput]{},
		).
		Put("/{id}", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[UpdateWorkScheduleInput](core.InBody, ctx)

			model := input.Dto()
			err := svc.Update(*contextInfo, ctx.Path("id"), model)
			if err != nil {
				return err
			}

			return ctx.JSON(core.Map{
				"status": true,
			})
		})

	ctrl.
		Pipe(core.PathParser[shared.ParamID]{}).
		Delete("/{id}", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)

			err := svc.Delete(*contextInfo, ctx.Path("id"))
			if err != nil {
				return err
			}

			return ctx.JSON(core.Map{
				"status": true,
			})
		})

	return ctrl
}
