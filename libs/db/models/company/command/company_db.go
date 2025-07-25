package command

import (
	tenantCommand "github.com/hros-aio/apis/libs/db/models/tenant/command"
	"github.com/tinh-tinh/sqlorm/v2"
)

type CompanyDB struct {
	sqlorm.Model     `gorm:"embedded"`
	TenantId         string                        `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_tenant_id,unique" json:"tenantId"`
	Name             string                        `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Industry         string                        `gorm:"column:industry;type:varchar(64);not null;" json:"industry"`
	Size             int                           `gorm:"column:size;type:int;not null;" json:"size"`
	Logo             string                        `gorm:"column:logo;type:varchar(256);not null;" json:"logo"`
	Contact          tenantCommand.ContactPersonDb `gorm:"embedded;embeddedPrefix:contact_" json:"contact"`
	SecondaryContact tenantCommand.ContactPersonDb `gorm:"embedded;embeddedPrefix:secondary_contact_" json:"secondaryContact"`
}

func (CompanyDB) TableName() string {
	return "companies"
}
