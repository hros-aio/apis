package roles

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/apps/auth-svc/app/permissions"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
)

type RoleDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantID     string                     `gorm:"type:varchar(255);not null;index:idx_roles_tenant_id"`
	CompanyID    uuid.UUID                  `gorm:"column:company_id" json:"companyId"`
	Company      *company.CompanyDB         `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Code         string                     `gorm:"column:code;type:varchar(64);not null" json:"code"`
	Name         string                     `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description  string                     `gorm:"type:text;" json:"description"`
	Permissions  []permissions.PermissionDB `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

func (RoleDB) TableName() string {
	return "roles"
}

func (data RoleDB) Dto() *RoleModel {
	model := &RoleModel{
		Model:       base.Model(data.Model),
		TenantID:    data.TenantID,
		CompanyID:   data.CompanyID,
		Code:        data.Code,
		Name:        data.Name,
		Description: data.Description,
	}

	if data.Company != nil {
		model.Company = data.Company.Dto()
	}

	if len(data.Permissions) > 0 {
		model.Permissions = make([]*permissions.PermissionModel, len(data.Permissions))
		for i, perm := range data.Permissions {
			model.Permissions[i] = perm.Dto()
		}
	}

	return model
}
