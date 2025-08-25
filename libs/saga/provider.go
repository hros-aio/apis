package saga

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/cacher/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

const EVENT_PUBLISHER = "EventPublisher"

type EventPulisher struct {
	clientProxy microservices.ClientProxy
	logger      *logger.Logger
	cacheSaga   *cacher.Schema[messages.SyncRegisteredPayload]
}

func NewProvider(module core.Module) core.Provider {
	client := microservices.InjectClient(module, microservices.KAFKA)
	logger := logger.InjectLog(module)
	cacheSaga := cacher.InjectSchemaByStore[messages.SyncRegisteredPayload](module, cacher.MEMORY)

	return module.NewProvider(core.ProviderOptions{
		Name: EVENT_PUBLISHER,
		Value: &EventPulisher{
			clientProxy: client,
			logger:      logger,
			cacheSaga:   cacheSaga,
		},
	})
}

func (p *EventPulisher) Publish(eventType string, data any, headers ...microservices.Header) {
	err := p.clientProxy.Publish(eventType, data, headers...)
	if err != nil {
		p.logger.Errorf("Failed to publish %v to event %s, because is %s", data, eventType, err.Error())
		return
	}
	p.logger.Infof("Publish message %v to event %s successfully", data, eventType)
}

func (p *EventPulisher) RegisterSync(eventType string, data messages.SyncDataPayload) error {
	sessionId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	payload := messages.SyncRegisteredPayload{
		SessionId: sessionId,
		Event:     eventType,
		SyncDataPayload: messages.SyncDataPayload{
			PreviousData: data.PreviousData,
			Data:         data.Data,
		},
	}

	err = p.cacheSaga.Set(sessionId.String(), payload)
	if err != nil {
		return err
	}

	return p.clientProxy.Publish(events.SyncRegistered, payload)
}

func (p *EventPulisher) RetrySync(sessionId string) error {
	payload, err := p.cacheSaga.Get(sessionId)
	if err != nil {
		return err
	}

	return p.clientProxy.Publish(events.SyncRetry, payload)
}
