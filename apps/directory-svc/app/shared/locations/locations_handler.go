package locations

import (
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	locationRepo := module.Ref(location.REPOSITORY).(*location.Repository)
	log := logger.InjectLog(module)

	handler.OnEvent(events.LocationCreated, saga.SyncFnc(module, func(ctx microservices.Ctx) error {
		var data messages.LocationCreatedPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			log.Error("Failed to parse message", logger.Metadata{
				"err":       err,
				"eventType": events.LocationCreated,
			})
			return err
		}

		model := ToModel(data)
		_, err = locationRepo.Create(model)
		if err != nil {
			log.Error("Failed to create location", logger.Metadata{
				"err":   err,
				"input": model,
			})
			return err
		}

		return nil
	}))

	return handler
}
