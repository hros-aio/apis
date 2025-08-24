package sync

import (
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/queue/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func NewProcessor(module core.Module) core.Provider {
	processor := queue.NewProcessor(QUEUE_NAME, module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	processor.Process(func(job *queue.Job) {
		sessionId := job.Id
		data, ok := job.Data.(messages.SyncRegisteredPayload)
		if !ok {
			return
		}

		eventPublisher.Publish(data.Event, data.SyncDataPayload, microservices.Header{
			"sessionId": sessionId,
		})
	})
	return processor
}
