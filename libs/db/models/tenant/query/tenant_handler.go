package query

import (
	"fmt"

	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	model := mongoose.InjectModel[TenantSchema](module)

	var tenant TenantSchema
	handler.OnEvent(tenant.CollectionName(), func(ctx microservices.Ctx) error {
		var input TenantSchema
		if err := ctx.PayloadParser(&input); err != nil {
			return err
		}
		action := ctx.Headers("action")
		switch action {
		case "create":
			_, err := model.Create(&input)
			if err != nil {
				return err
			}
			return nil
		case "update":
			err := model.UpdateByID(input.ID, &input)
			if err != nil {
				return err
			}
			return nil
		case "delete":
			err := model.DeleteByID(input.ID)
			if err != nil {
				return err
			}
			return nil
		default:
			return fmt.Errorf("unknown action: %s", action)
		}
	})

	return handler
}
