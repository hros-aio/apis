package company

import (
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
)

type CompanyModel struct {
	base.Model
	TenantId         string               `json:"tenantId"`
	Name             string               `json:"name"`
	Industry         string               `json:"industry"`
	Size             int                  `json:"size"`
	Logo             string               `json:"logo"`
	Contact          tenant.ContactPerson `json:"contact"`
	SecondaryContact tenant.ContactPerson `json:"secondaryContact"`
}
