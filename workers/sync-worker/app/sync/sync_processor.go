package sync

import (
	"fmt"

	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/queue/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewProcessor(module core.Module) core.Provider {
	processor := queue.NewProcessor(QUEUE_NAME, module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)
	log := logger.InjectLog(module)

	processor.Process(func(job *queue.Job) {
		sessionId := job.Id
		data, ok := job.Data.(messages.SyncRegisteredPayload)
		if !ok {
			log.Error("Failed to cast job data", logger.Metadata{
				"error":   "Invalid job data",
				"payload": fmt.Sprintf("%+v", job.Data),
			})
			return
		}

		eventPublisher.Publish(data.Event, data.Data, microservices.Header{
			"sessionId": sessionId,
		})
	})
	return processor
}
