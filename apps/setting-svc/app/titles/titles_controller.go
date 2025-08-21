package titles

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/factory/shared"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	svc := core.Inject[TitleService](module)
	ctrl := module.NewController("titles").
		Metadata(swagger.ApiTag("Title"), swagger.ApiSecurity("bearerAuth")).
		Registry()

	ctrl.
		Pipe(
			core.QueryParser[shared.QueryCompany]{},
			core.BodyParser[CreateTitleInput]{},
		).
		Post("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[CreateTitleInput](core.InBody, ctx)

			model := input.Dto(*contextInfo)
			data, err := svc.Create(*contextInfo, model)
			if err != nil {
				return err
			}

			return ctx.JSON(core.Map{
				"data": data,
			})
		})

	ctrl.
		Use(middleware.Pagination).
		Get("", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			queryParams := core.Execution[middleware.Paginate](core.InQuery, ctx)

			data, total, err := svc.List(*contextInfo, *queryParams)
			if err != nil {
				return err
			}

			return ctx.JSON(core.Map{
				"data":  data,
				"total": total,
			})
		})

	ctrl.
		Pipe(
			core.PathParser[shared.ParamID]{},
			core.BodyParser[UpdateTitleInput]{},
		).
		Put("{id}", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[UpdateTitleInput](core.InBody, ctx)

			model := input.Dto()
			data, err := svc.UpdateByID(*contextInfo, ctx.Path("id"), model)
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

			err := svc.DeleteByID(*contextInfo, ctx.Path("id"))
			if err != nil {
				return err
			}

			return ctx.JSON(core.Map{
				"status": true,
			})
		})

	return ctrl
}
