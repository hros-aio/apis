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
	Line     string `json:"addressLine" validate:"required" example:"123 Main St"`
	City     string `json:"city" validate:"required" example:"Anytown"`
	State    string `json:"state" validate:"required" example:"CA"`
	Country  string `json:"country" validate:"required" example:"USA"`
	Zipcode  string `json:"zipCode" validate:"required" example:"12345"`
	Timezone string `json:"timezone" validate:"required" example:"America/Los_Angeles"`
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
