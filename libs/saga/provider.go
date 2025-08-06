package saga

import (
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

const EVENT_PUBLISHER = "EventPublisher"

type EventPulisher struct {
	clientProxy microservices.ClientProxy
	logger      *logger.Logger
}

func NewProvider(module core.Module) core.Provider {
	client := microservices.InjectClient(module, microservices.KAFKA)
	logger := logger.InjectLog(module)

	return module.NewProvider(core.ProviderOptions{
		Name: EVENT_PUBLISHER,
		Value: &EventPulisher{
			clientProxy: client,
			logger:      logger,
		},
	})
}

func (p *EventPulisher) Publish(eventType string, data any) {
	err := p.clientProxy.Publish(eventType, data)
	if err != nil {
		p.logger.Errorf("Failed to publish %v to event %s, because is %s", data, eventType, err.Error())
		return
	}
	p.logger.Infof("Publish message %v to event %s successfully", data, eventType)
}
