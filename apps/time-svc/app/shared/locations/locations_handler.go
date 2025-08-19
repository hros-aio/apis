package locations

import (
	"time-svc/common/constants"

	"github.com/hros-aio/apis/libs/mongodoc/common/location"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/pubsub/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	locationRepo := module.Ref(location.REPOSITORY).(*location.Repository)
	logger := logger.InjectLog(module)
	pubsub := pubsub.InjectBroker(module)

	handler.OnEvent(events.LocationCreated, func(ctx microservices.Ctx) error {
		var data messages.LocationCreatedPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		model := ToModel(data)
		result, err := locationRepo.Create(model)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		go pubsub.Publish(constants.EVENT_INTERNAL_LOCATION_CREATED, result)
		return nil
	})

	return handler
}
