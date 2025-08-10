package auth

import (
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/common/exception"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("auth").
		Metadata(swagger.ApiTag("Auth")).
		Registry()

	svc := core.Inject[AuthService](module)
	ctrl.
		Pipe(core.BodyParser[LoginInput]{}).
		Post("login", func(ctx core.Ctx) error {
			contextInfo, ok := ctx.Get(middleware.APP_CONTEXT).(middleware.ContextInfo)
			if !ok {
				return exception.InternalServer("Cannot parse context")
			}
			input, ok := ctx.Body().(*LoginInput)
			if !ok {
				return exception.BadRequest("Cannot parse input")
			}
			data, err := svc.Login(contextInfo, input)
			if err != nil {
				return err
			}
			return ctx.JSON(data)
		})

	return ctrl
}
