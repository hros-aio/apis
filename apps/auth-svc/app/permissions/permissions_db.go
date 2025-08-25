package permissions

import (
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/tinh-tinh/sqlorm/v2"
)

type PermissionDB struct {
	sqlorm.Model `gorm:"embedded"`
	Name         string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Module       string `gorm:"type:varchar(255);not null;index"`
}

func (PermissionDB) TableName() string {
	return "permissions"
}

func (data PermissionDB) Dto() *PermissionModel {
	return &PermissionModel{
		Model:  base.Model(data.Model),
		Name:   data.Name,
		Module: data.Module,
	}
}
