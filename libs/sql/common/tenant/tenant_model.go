package tenant

import "github.com/hros-aio/apis/libs/sql/common/base"

type ContactPerson struct {
	ContactName  string `json:"name" example:"Renil"`
	ContactEmail string ` json:"email" example:"renil@gmail.com" validate:"isEmail"`
	ContactPhone string `json:"phone" example:"012345678"`
}

type TenantModel struct {
	base.Model
	TenantId    string        `json:"tenantId"`
	Name        string        `json:"name"`
	Description string        ` json:"description"`
	Contact     ContactPerson `json:"contact"`
}
