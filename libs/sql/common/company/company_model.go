package company

import (
	"github.com/hros-aio/apis/libs/sql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
)

type CompanyModel struct {
	sqlorm.Model     `gorm:"embedded"`
	TenantId         string                 `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_tenant_id,unique" json:"tenantId"`
	Name             string                 `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Industry         string                 `gorm:"column:industry;type:varchar(64);not null;" json:"industry"`
	Size             int                    `gorm:"column:size;type:int;not null;" json:"size"`
	Logo             string                 `gorm:"column:logo;type:varchar(256);not null;" json:"logo"`
	Contact          tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:contact_" json:"contact"`
	SecondaryContact tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:secondary_contact_" json:"secondaryContact"`
}

func (CompanyModel) TableName() string {
	return "companies"
}
