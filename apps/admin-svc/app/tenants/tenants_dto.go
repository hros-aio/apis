package tenants

import (
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/hros-aio/apis/libs/saga/messages"
)

type ContactInput struct {
	Name  string `json:"name" example:"Abc"`
	Email string `json:"email" example:"abc@gmail.com" validate:"isEmail"`
	Phone string `json:"phone" example:"012345678"`
}

type TenantCreateInput struct {
	Name        string        `json:"name" example:"Abc Inc" validate:"required"`
	Domain      string        `json:"domain" example:"abc.hros.aio" validate:"required"`
	Description string        `json:"description" example:"Abc Inc" validate:"required"`
	Contact     *ContactInput `json:"contact" validate:"nested"`
}

func (m *TenantCreateInput) Dto() *tenant.TenantDB {
	return &tenant.TenantDB{
		Name:        m.Name,
		Domain:      m.Domain,
		Description: m.Description,
		Contact: tenant.ContactPersonDb{
			ContactName:  m.Contact.Name,
			ContactEmail: m.Contact.Email,
			ContactPhone: m.Contact.Phone,
		},
	}
}

func ToMessage(data *tenant.TenantModel) messages.TenantCreatedPayload {
	return messages.TenantCreatedPayload{
		Id:        data.ID.String(),
		Name:      data.Contact.ContactName,
		CreatedAt: data.CreatedAt,
		TenantId:  data.TenantId,
		Domain:    data.Domain,
		Contact:   messages.ContactPerson(data.Contact),
	}
}
