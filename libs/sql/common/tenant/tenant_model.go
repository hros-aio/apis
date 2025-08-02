package tenant

import "github.com/hros-aio/apis/libs/sql/common/base"

type ContactPerson struct {
	ContactName  string `json:"name"`
	ContactEmail string ` json:"email"`
	ContactPhone string `json:"phone"`
}

type TenantModel struct {
	base.Model
	TenantId    string        `json:"tenantId"`
	Name        string        `json:"name"`
	Description string        ` json:"description"`
	Contact     ContactPerson `json:"contact"`
}
