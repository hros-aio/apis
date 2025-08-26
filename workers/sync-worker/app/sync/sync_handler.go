package sync

import (
	"fmt"

	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/queue/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	cacheSync := cacher.InjectSchemaByStore[messages.SyncRegisteredPayload](module, cacher.MEMORY)
	syncQueue := queue.Inject(module, QUEUE_NAME)
	log := logger.InjectLog(module)

	handler.OnEvent(events.SyncRegistered, func(ctx microservices.Ctx) error {
		var payload messages.SyncRegisteredPayload
		if err := ctx.PayloadParser(&payload); err != nil {
			log.Error("Failed to parse payload", logger.Metadata{
				"error":   err.Error(),
				"payload": fmt.Sprintf("%+v", payload),
			})
			return err
		}

		err := cacheSync.Set(payload.SessionId.String(), payload)
		if err != nil {
			log.Error("Failed to cache payload", logger.Metadata{
				"error":   err.Error(),
				"payload": fmt.Sprintf("%+v", payload),
			})
			return err
		}

		syncQueue.AddJob(queue.AddJobOptions{
			Id:   payload.SessionId.String(),
			Data: payload,
		})
		return nil
	})

	handler.OnEvent(events.SyncRetry, func(ctx microservices.Ctx) error {
		var payload messages.SyncRetryPayload
		if err := ctx.PayloadParser(&payload); err != nil {
			log.Error("Failed to parse payload", logger.Metadata{
				"error":   err.Error(),
				"payload": fmt.Sprintf("%+v", payload),
			})
			return err
		}

		cached, err := cacheSync.Get(payload.SessionId.String())
		if err != nil {
			log.Error("Failed to get cached payload", logger.Metadata{
				"error":   err.Error(),
				"payload": fmt.Sprintf("%+v", payload),
			})
			return err
		}

		syncQueue.AddJob(queue.AddJobOptions{
			Id:   payload.SessionId.String(),
			Data: cached,
		})

		return nil
	})

	return handler
}
