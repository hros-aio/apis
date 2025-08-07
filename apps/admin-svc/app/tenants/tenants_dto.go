package tenants

import "github.com/hros-aio/apis/libs/psql/common/tenant"

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
