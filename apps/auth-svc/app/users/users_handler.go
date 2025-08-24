package users

import (
	"strings"

	"github.com/hros-aio/apis/libs/psql/common/user"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/auth/v2"
	"github.com/tinh-tinh/tinhtinh/microservices"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

func NewHandler(module core.Module) core.Provider {
	handler := microservices.NewHandler(module, core.ProviderOptions{})

	repo := module.Ref(user.REPOSITORY).(*user.Repository)
	logger := logger.InjectLog(module)

	handler.OnEvent(events.TenantCreated, func(ctx microservices.Ctx) error {
		var data messages.TenantCreatedPayload
		err := ctx.PayloadParser(&data)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		username := strings.Split(data.Contact.ContactEmail, "@")
		model := &user.UserDB{
			TenantId:   data.TenantId,
			Username:   username[0],
			Email:      data.Contact.ContactEmail,
			Password:   auth.Hash("12345678@Tc"),
			IsVerified: false,
			IsBanned:   false,
			IsAdmin:    true,
		}
		createdUser, err := repo.Create(model)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		logger.Infof("Create user successfully: %v", createdUser)
		return nil
	})

	return handler
}
