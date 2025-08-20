package titles

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/apps/setting-svc/app/grades"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/tinh-tinh/sqlorm/v2"
)

type TitleModel struct {
	base.Model
	TenantID     string                      `json:"tenantId"`
	CompanyID    uuid.UUID                   `json:"companyId"`
	Company      *company.CompanyModel       `json:"company,omitempty"`
	DepartmentID uuid.UUID                   `json:"departmentId"`
	Department   *department.DepartmentModel `json:"department,omitempty"`
	GradeID      uuid.UUID                   `json:"gradeId"`
	Grade        *grades.GradeModel          `json:"grade,omitempty"`
	Code         string                      `json:"code"`
	Name         string                      `json:"name"`
}

func (model TitleModel) DataMapper() *TitleDB {
	data := &TitleDB{
		TenantID:     model.TenantID,
		CompanyID:    model.CompanyID,
		DepartmentID: model.DepartmentID,
		GradeID:      model.GradeID,
		Code:         model.Code,
		Name:         model.Name,
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}

	return data
}
