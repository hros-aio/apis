package department

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
)

type DepartmentModel struct {
	base.Model
	TenantID   string                `json:"tenantId"`
	CompanyID  uuid.UUID             `json:"companyId"`
	Company    *company.CompanyModel `json:"company,omitempty"`
	Code       string                `json:"code"`
	Name       string                `json:"name"`
	IsDivision bool                  `json:"isDivision"`
	ParentID   *uuid.UUID            `json:"parentId"`
	Parent     *DepartmentModel      `json:"parent,omitempty"`
}

func (model DepartmentModel) DataMapper() *DepartmentDB {
	data := &DepartmentDB{
		TenantID:   model.TenantID,
		Name:       model.Name,
		Code:       model.Code,
		IsDivision: model.IsDivision,
		CompanyID:  model.CompanyID,
		ParentID:   model.ParentID,
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}

	return data
}
