package snapshot

import (
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	snapshotRepo := mongoose.InjectModel[SnapshotSchema](module)
	logger := logger.InjectLog(module)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)

	handler.OnEvent(events.SagaRegistered, func(ctx microservices.Ctx) error {
		// Payload parse message
		var data messages.SagaRegisteredPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		model := &SnapshotSchema{
			SessionId: data.SessionId.String(),
			Event:     data.Event,
			Step:      data.Step,
		}

		// Convert data any to bson M
		bsonBytes, err := bson.Marshal(data.Data)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		var bm bson.M
		if err := bson.Unmarshal(bsonBytes, &bm); err != nil {
			logger.Error(err.Error())
			return err
		}

		// Create snapshot
		model.Payload = bm
		_, err = snapshotRepo.Create(model)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		go eventPublisher.Publish(data.Event, data.Data)
		return nil
	})

	handler.OnEvent(events.SagaRollback, func(ctx microservices.Ctx) error {
		// Payload parse message
		var data messages.SagaRollbackPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		snapshots, err := snapshotRepo.Find(bson.M{
			"sessionId": data.SessionId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		for _, snapshot := range snapshots {
			go eventPublisher.Publish(snapshot.Event, snapshot.Payload)
		}

		return nil
	})

	return handler
}
