package tenant

import "github.com/tinh-tinh/sqlorm/v2"

type CommandDb struct {
	sqlorm.Model `gorm:"embedded"`
	TenantId     string `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_tenant_id,unique" json:"tenant_id"`
	Name         string `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Description  string `gorm:"column:description;type:varchar(256);not null;" json:"description"`
}

func (CommandDb) TableName() string {
	return "tenants"
}
