package tenants

import (
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/hros-aio/apis/libs/saga/messages"
)

type ContactInput struct {
	Name  string `json:"name" example:"Renil"`
	Email string `json:"email" example:"renil@gmail.com" validate:"isEmail"`
	Phone string `json:"phone" example:"012345678"`
}

type TenantCreateInput struct {
	Name        string        `json:"name" example:"Netze" validate:"required"`
	Domain      string        `json:"domain" example:"netze.home.ai" validate:"required"`
	Description string        `json:"description" example:"Netze Home Inc" validate:"required"`
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
