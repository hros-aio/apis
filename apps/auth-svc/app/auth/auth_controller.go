package auth

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	svc := core.Inject[AuthService](module)
	ctrl := module.NewController("auth").
		Metadata(swagger.ApiTag("Auth")).
		Registry()

	ctrl.
		Metadata(middleware.IsPublic()).
		Pipe(core.BodyParser[LoginInput]{}).
		Post("login", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[LoginInput](core.InBody, ctx)
			data, err := svc.Login(contextInfo, input)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	ctrl.
		Metadata(swagger.ApiSecurity("bearerAuth")).
		Get("me", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			data, err := svc.GetMe(contextInfo)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	ctrl.
		Metadata(middleware.IsPublic()).
		Pipe(core.BodyParser[RefreshAccessTokenInput]{}).
		Post("refresh-access-token", func(ctx core.Ctx) error {
			contextInfo := core.Execution[middleware.ContextInfo](middleware.APP_CONTEXT, ctx)
			input := core.Execution[RefreshAccessTokenInput](core.InBody, ctx)
			data, err := svc.RefreshAccessToken(contextInfo, input)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	return ctrl
}
