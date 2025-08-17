package location

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
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
	TenantId      string                `json:"tenantId"`
	CompanyID     uuid.UUID             `json:"companyId"`
	Company       *company.CompanyModel `json:"company,omitempty"`
	Name          string                `json:"name"`
	Contact       tenant.ContactPerson  `json:"contact"`
	MapUrl        string                `json:"mapUrl"`
	IsHeadquarter bool                  `json:"isHeadquarter"`
}

func (model LocationModel) DataMapper() *LocationDB {
	data := &LocationDB{
		TenantId:      model.TenantId,
		Name:          model.Name,
		AddressInfoDB: AddressInfoDB(model.AddressInfo),
		CompanyID:     model.CompanyID,
		Contact:       tenant.ContactPersonDb(model.Contact),
		MapUrl:        model.MapUrl,
		IsHeadquarter: model.IsHeadquarter,
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}

	return data
}
