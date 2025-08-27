package departments

import (
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	departmentRepo := module.Ref(department.REPOSITORY).(*department.Repository)
	log := logger.InjectLog(module)

	handler.OnEvent(events.DepartmentCreated, saga.SyncFnc(module, func(ctx microservices.Ctx) error {
		var data messages.DepartmentCreatedPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			log.Error("Failed to parse department created payload", logger.Metadata{
				"err": err.Error(),
			})
			return err
		}

		model := ToModel(data)
		_, err = departmentRepo.Create(model)
		if err != nil {
			log.Error("Failed to create department", logger.Metadata{
				"err": err.Error(),
			})
			return err
		}
		return nil
	}))

	return handler
}
