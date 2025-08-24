package location

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
)

type AddressInfoDB struct {
	Line     string `gorm:"column:address_line;type:varchar(255);not null" json:"addressLine"`
	City     string `gorm:"column:city;type:varchar(255);not null" json:"city"`
	State    string `gorm:"column:state;type:varchar(100);not null" json:"state"`
	Country  string `gorm:"column:country;type:varchar(100);not null" json:"country"`
	Zipcode  string `gorm:"column:zip_code;type:varchar(64);not null" json:"zipCode"`
	Timezone string `gorm:"column:timezone;type:varchar(64);not null" json:"timezone"`
}

type LocationDB struct {
	sqlorm.Model  `gorm:"embedded"`
	AddressInfoDB `gorm:"embedded"`
	TenantId      string                 `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_location_tenant_id" json:"tenantId"`
	CompanyID     uuid.UUID              `gorm:"column:company_id" json:"companyId"`
	Company       *company.CompanyDB     `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Name          string                 `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Contact       tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:contact_" json:"contact"`
	MapUrl        string                 `gorm:"column:map_url" json:"mapUrl"`
	IsHeadquarter bool                   `gorm:"is_headquarter;default:false" json:"isHeadquarter"`
}

func (LocationDB) TableName() string {
	return "locations"
}

func (data LocationDB) Dto() *LocationModel {
	model := &LocationModel{
		Model:         base.Model(data.Model),
		AddressInfo:   AddressInfo(data.AddressInfoDB),
		TenantId:      data.TenantId,
		CompanyID:     data.CompanyID,
		Name:          data.Name,
		Contact:       tenant.ContactPerson(data.Contact),
		MapUrl:        data.MapUrl,
		IsHeadquarter: data.IsHeadquarter,
	}

	if data.Company != nil {
		model.Company = data.Company.Dto()
	}

	return model
}
