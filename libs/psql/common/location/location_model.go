package location

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
)

type AddressInfo struct {
	Line     string `json:"addressLine"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zipcode  string `json:"zipCode"`
	Timezone string `json:"timezone"`
}

type LocationModel struct {
	base.Model
	AddressInfo
	TenantId  string                `json:"tenantId"`
	CompanyID uuid.UUID             `json:"companyId"`
	Company   *company.CompanyModel `json:"company,omitempty"`
	Name      string                `json:"name"`
	Contact   tenant.ContactPerson  `json:"contact"`
	MapUrl    string                `json:"mapUrl"`
}
