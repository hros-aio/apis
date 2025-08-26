package employees

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/employee"
)

type CreateEmployeeInput struct {
	GeneralInfo  *employee.GeneralInfo `json:"generalInfo" validate:"nested"`
	LocationID   string                `json:"locationId" validate:"required,isUUID" example:"0e9b0ae1-dce9-4e25-9cf2-a669738d1f21"`
	DepartmentID string                `json:"departmentId" validate:"required,isUUID" example:"0e9b0ae1-dce9-4e25-9cf2-a669738d1f21"`
	GradeID      string                `json:"gradeId" validate:"required,isUUID" example:"0e9b0ae1-dce9-4e25-9cf2-a669738d1f21"`
	TitleID      string                `json:"titleId" validate:"required,isUUID" example:"0e9b0ae1-dce9-4e25-9cf2-a669738d1f21"`
	JoiningDate  time.Time             `json:"joiningDate" validate:"required,isDate" example:"2023-01-01"`
	Code         string                `json:"code" validate:"required" example:"EMP001"`
	Type         string                `json:"type" validate:"required" example:"full-time"`
}

func (data CreateEmployeeInput) Dto(ctx middleware.ContextInfo) *employee.EmployeeModel {
	names := []string{}
	if data.GeneralInfo != nil {
		names = []string{data.GeneralInfo.FirstName, data.GeneralInfo.MiddleName, data.GeneralInfo.LastName}
	}
	return &employee.EmployeeModel{
		TenantID:     ctx.TenantID,
		CompanyID:    ctx.CompanyID,
		GeneralInfo:  *data.GeneralInfo,
		LocationID:   uuid.MustParse(data.LocationID),
		DepartmentID: uuid.MustParse(data.DepartmentID),
		GradeID:      uuid.MustParse(data.GradeID),
		TitleID:      uuid.MustParse(data.TitleID),
		JoiningDate:  data.JoiningDate,
		Code:         data.Code,
		Type:         data.Type,
		Status:       employee.Employed,
		FullName:     strings.Join(names, " "),
	}
}
