package command

import (
	companyCommand "github.com/hros-aio/apis/libs/db/models/company/command"
	"github.com/tinh-tinh/sqlorm/v2"
)

type LocationDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantId     string                    `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_tenant_id,unique" json:"tenantId"`
	CompanyId    string                    `gorm:"column:company_id;type:varchar(64);not null;index:idx_company_id,unique" json:"companyId"`
	Company      *companyCommand.CompanyDB `gorm:"foreignKey:CompanyId;references:Id" json:"company"`
	Name         string                    `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	AddressLine  string                    `gorm:"column:address_line;type:varchar(256);not null;" json:"addressLine"`
	City         string                    `gorm:"column:city;type:varchar(64);not null;" json:"city"`
	State        string                    `gorm:"column:state;type:varchar(64);not null;" json:"state"`
	Country      string                    `gorm:"column:country;type:varchar(64);not null;" json:"country"`
	PostalCode   string                    `gorm:"column:postal_code;type:varchar(20);not null;" json:"postalCode"`
}

func (LocationDB) TableName() string {
	return "locations"
}
