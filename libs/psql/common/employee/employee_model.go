package employee

import (
	"time"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/hros-aio/apis/libs/psql/common/user"
)

type EmployeeModel struct {
	base.Model
	TenantID     string                      `json:"tenantId"`
	CompanyID    uuid.UUID                   `json:"companyId"`
	Company      *company.CompanyModel       `json:"company,omitempty"`
	UserID       uuid.UUID                   `json:"userId"`
	User         *user.UserModel             `json:"user,omitempty"`
	DepartmentID uuid.UUID                   `json:"departmentId"`
	Department   *department.DepartmentModel `json:"department,omitempty"`
	LocationID   uuid.UUID                   `json:"locationId"`
	Location     *location.LocationModel     `json:"location,omitempty"`
	GradeID      uuid.UUID                   `json:"gradeId"`
	TitleID      uuid.UUID                   `json:"titleId"`
	FullName     string                      `json:"fullName"`
	Code         string                      `json:"code"`
	JoiningDate  time.Time                   `json:"joiningDate"`
	LeftDate     *time.Time                  `json:"leftDate,omitempty"`
	Type         string                      `json:"type"`
	Status       string                      `json:"status"`
}

func (model EmployeeModel) DataMapper() *EmployeeDB {
	return &EmployeeDB{
		TenantID:     model.TenantID,
		CompanyID:    model.CompanyID,
		UserID:       model.UserID,
		DepartmentID: model.DepartmentID,
		LocationID:   model.LocationID,
		GradeID:      model.GradeID,
		TitleID:      model.TitleID,
		FullName:     model.FullName,
		Code:         model.Code,
		JoiningDate:  model.JoiningDate,
		LeftDate:     model.LeftDate,
		Type:         model.Type,
		Status:       model.Status,
	}
}
