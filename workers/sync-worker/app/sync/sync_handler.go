package sync

import (
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	cacheSaga := cacher.Inject[messages.SyncRegisteredPayload](module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	handler.OnEvent(events.SyncRegistered, func(ctx microservices.Ctx) error {
		var payload messages.SyncRegisteredPayload
		if err := ctx.PayloadParser(&payload); err != nil {
			return err
		}

		go eventPublisher.Publish(payload.Event, payload.SyncDataPayload, microservices.Header{
			"sessionId": payload.SessionId.String(),
		})
		return nil
	})

	handler.OnEvent(events.SyncRetry, func(ctx microservices.Ctx) error {
		var payload messages.SyncRetryPayload
		if err := ctx.PayloadParser(&payload); err != nil {
			return err
		}

		cached, err := cacheSaga.Get(payload.SessionId.String())
		if err != nil {
			return err
		}

		go eventPublisher.Publish(cached.Event, cached.SyncDataPayload, microservices.Header{
			"sessionId": cached.SessionId.String(),
		})
		return nil
	})

	return handler
}
