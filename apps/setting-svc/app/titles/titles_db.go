package titles

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/apps/setting-svc/app/grades"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/tinh-tinh/sqlorm/v2"
)

type TitleDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantID     string                   `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_location_tenant_id" json:"tenantId"`
	CompanyID    uuid.UUID                `gorm:"column:company_id" json:"companyId"`
	Company      *company.CompanyDB       `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	DepartmentID uuid.UUID                `gorm:"column:department_id" json:"departmentId"`
	Department   *department.DepartmentDB `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	GradeID      uuid.UUID                `gorm:"column:grade_id" json:"gradeId"`
	Grade        *grades.GradeDB          `gorm:"foreignKey:GradeID" json:"grade,omitempty"`
	Code         string                   `gorm:"column:code;type:varchar(64);not null" json:"code"`
	Name         string                   `gorm:"column:name;type:varchar(255)" json:"name"`
}

func (TitleDB) TableName() string {
	return "titles"
}

func (data TitleDB) Dto() *TitleModel {
	model := &TitleModel{
		Model:        base.Model(data.Model),
		TenantID:     data.TenantID,
		CompanyID:    data.CompanyID,
		DepartmentID: data.DepartmentID,
		GradeID:      data.GradeID,
		Code:         data.Code,
		Name:         data.Name,
	}

	if data.Company != nil {
		model.Company = data.Company.Dto()
	}
	if data.Department != nil {
		model.Department = data.Department.Dto()
	}
	if data.Grade != nil {
		model.Grade = data.Grade.Dto()
	}

	return model
}
