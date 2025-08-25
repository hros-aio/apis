package companies

import (
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	companyRepo := module.Ref(company.REPOSITORY).(*company.Repository)
	log := logger.InjectLog(module)

	handler.OnEvent(events.CompanyActivated, saga.SyncFnc(module, func(ctx microservices.Ctx) error {
		var data messages.CompanyActivatedPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			log.Error("Failed to parse message", logger.Metadata{
				"err":       err,
				"eventType": events.CompanyActivated,
			})
			return err
		}

		foundCompany, err := companyRepo.FindByID(data.ID)
		if err != nil {
			log.Error("Failed to find company", logger.Metadata{
				"err":       err,
				"eventType": events.CompanyActivated,
			})
			return err
		}

		model := ToModel(data)
		if foundCompany == nil {
			_, err = companyRepo.Create(model)
			if err != nil {
				log.Error("Failed to create company", logger.Metadata{
					"err":       err,
					"eventType": events.CompanyActivated,
				})
				return err
			}

			return nil
		}

		_, err = companyRepo.UpdateByID(data.ID, model)
		if err != nil {
			log.Error("Failed to update company", logger.Metadata{
				"err":       err,
				"eventType": events.CompanyActivated,
			})
			return err
		}

		return nil
	}))

	return handler
}
