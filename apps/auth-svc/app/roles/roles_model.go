package roles

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/apps/auth-svc/app/permissions"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
)

type RoleModel struct {
	base.Model
	TenantID    string                         `json:"tenantId"`
	CompanyID   uuid.UUID                      `json:"companyId"`
	Company     *company.CompanyModel          `json:"company,omitempty"`
	Code        string                         `json:"code"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	Permissions []*permissions.PermissionModel `json:"permissions,omitempty"`
}

func (model RoleModel) DataMapper() *RoleDB {
	data := &RoleDB{
		TenantID:    model.TenantID,
		CompanyID:   model.CompanyID,
		Code:        model.Code,
		Name:        model.Name,
		Description: model.Description,
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}
	return data
}
