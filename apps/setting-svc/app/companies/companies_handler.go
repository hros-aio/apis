package companies

import (
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})
	repo := module.Ref(company.REPOSITORY).(*company.Repository)
	logger := logger.InjectLog(module)

	handler.OnEvent(events.TenantCreated, func(ctx microservices.Ctx) error {
		var data messages.TenantCreatedPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		input := &company.CompanyModel{
			TenantID: data.TenantId,
			Name:     data.Name,
			Contact:  tenant.ContactPerson(data.Contact),
		}
		createdCompany, err := repo.Create(input)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		logger.Infof("Create company successfully: %v", createdCompany)
		return nil
	})

	return handler
}
