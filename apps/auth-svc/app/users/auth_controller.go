package users

import (
	"github.com/tinh-tinh/swagger/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func AuthController(module core.Module) core.Controller {
	ctrl := module.NewController("auth").
		Metadata(swagger.ApiTag("Auth")).
		Registry()

	return ctrl
}
