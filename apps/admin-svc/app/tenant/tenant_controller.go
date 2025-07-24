package tenant

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("tenants")

	return ctrl
}
