package company

import (
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
)

type CompanyDB struct {
	sqlorm.Model     `gorm:"embedded"`
	TenantId         string                 `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_company_tenant_id" json:"tenantId"`
	Name             string                 `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Industry         string                 `gorm:"column:industry;type:varchar(64);not null;" json:"industry"`
	Size             int                    `gorm:"column:size;type:int;not null;" json:"size"`
	Logo             string                 `gorm:"column:logo;type:varchar(256);not null;" json:"logo"`
	Contact          tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:contact_" json:"contact"`
	SecondaryContact tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:secondary_contact_" json:"secondaryContact"`
}

func (CompanyDB) TableName() string {
	return "companies"
}

func (data CompanyDB) Dto() *CompanyModel {
	return &CompanyModel{
		Model:            base.Model(data.Model),
		TenantId:         data.TenantId,
		Name:             data.Name,
		Industry:         data.Industry,
		Size:             data.Size,
		Logo:             data.Logo,
		Contact:          tenant.ContactPerson(data.Contact),
		SecondaryContact: tenant.ContactPerson(data.SecondaryContact),
	}
}
