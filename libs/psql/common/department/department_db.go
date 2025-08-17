package department

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
)

type DepartmentDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantID     string             `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_location_tenant_id" json:"tenantId"`
	CompanyID    uuid.UUID          `gorm:"column:company_id" json:"companyId"`
	Company      *company.CompanyDB `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Code         string             `gorm:"column:code;type:varchar(64);not null" json:"code"`
	Name         string             `gorm:"column:name;type:varchar(255)" json:"name"`
	IsDivision   bool               `gorm:"column:is_division;default:false" json:"isDivision"`
	ParentID     *uuid.UUID         `gorm:"column:parent_id" json:"parentId"`
	Parent       *DepartmentDB      `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
}

func (DepartmentDB) TableName() string {
	return "departments"
}

func (data DepartmentDB) Dto() *DepartmentModel {
	model := &DepartmentModel{
		Model:      base.Model(data.Model),
		TenantID:   data.TenantID,
		CompanyID:  data.CompanyID,
		Code:       data.Code,
		Name:       data.Name,
		IsDivision: data.IsDivision,
		ParentID:   data.ParentID,
	}

	if data.Company != nil {
		model.Company = data.Company.Dto()
	}

	if data.Parent != nil {
		model.Parent = data.Parent.Dto()
	}

	return model
}
