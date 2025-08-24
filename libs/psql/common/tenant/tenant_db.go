package tenant

import "github.com/tinh-tinh/sqlorm/v2"

type ContactPersonDb struct {
	ContactName  string `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	ContactEmail string `gorm:"column:email;type:varchar(256);not null;" json:"email"`
	ContactPhone string `gorm:"column:phone;type:varchar(64);not null;" json:"phone"`
}

type TenantDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantId     string          `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_tenant_id,unique" json:"tenantId"`
	Domain       string          `gorm:"column:domain;type:varchar(64);not null;index:idx_tenant_domain,unique" json:"domain"`
	Name         string          `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Description  string          `gorm:"column:description;type:varchar(256);not null;" json:"description"`
	Contact      ContactPersonDb `gorm:"embedded;embeddedPrefix:contact_" json:"contact"`
}

func (TenantDB) TableName() string {
	return "tenants"
}
