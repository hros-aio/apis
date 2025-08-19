package works_chedules

import (
	"time-svc/common/constants"

	"github.com/hros-aio/apis/libs/mongodoc/common/location"
	"github.com/tinh-tinh/mongoose/v2"
	"github.com/tinh-tinh/pubsub/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func InternalHandler(module core.Module) core.Provider {
	handler := pubsub.NewHandler(module)
	model := mongoose.InjectModel[WorkScheduleSchema](module)
	log := logger.InjectLog(module)

	handler.Listen(func(sub *pubsub.Message) {
		msg := sub.GetContent()
		result, ok := msg.(*location.LocationModel)
		if !ok {
			log.Error("Failed to cast message to InsertOneResult")
			return
		}

		// Handle the result
		data := &WorkScheduleSchema{
			TenantID:   result.TenantID,
			CompanyID:  result.CompanyID,
			LocationID: result.ID,
			Name:       "Default",
			IsDefault:  true,
			TotalHours: constants.TotalHoursDefault,
			StartAt:    constants.StartAtDefault,
			EndAt:      constants.EndAtDefault,
			WorkDays:   []int{1, 2, 3, 4, 5},
		}
		success, err := model.Create(data)
		if err != nil {
			log.Error("Failed to create work schedule:", logger.Metadata{
				"err": err,
			})
			return
		}
		log.Info("Work schedule created successfully", logger.Metadata{
			"id": success.InsertedID,
		})
	}, constants.EVENT_INTERNAL_LOCATION_CREATED)

	return handler
}
