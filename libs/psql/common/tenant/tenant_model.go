package tenant

import "github.com/hros-aio/apis/libs/psql/common/base"

type ContactPerson struct {
	ContactName  string `json:"name" example:"Abc"`
	ContactEmail string `json:"email" example:"abc@gmail.com" validate:"isEmail"`
	ContactPhone string `json:"phone" example:"012345678"`
}

type TenantModel struct {
	base.Model
	TenantId    string        `json:"tenantId"`
	Domain      string        `json:"domain"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Contact     ContactPerson `json:"contact"`
}
