package employee

import (
	"time"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/psql/common/location"
)

type GeneralInfo struct {
	FirstName     string               `json:"firstName"`
	MiddleName    string               `json:"middleName"`
	LastName      string               `json:"lastName"`
	Birth         string               `json:"birth"`
	Gender        string               `json:"gender"`
	WorkEmail     string               `json:"workEmail"`
	WorkPhone     string               `json:"workPhone"`
	PersonalEmail string               `json:"personalEmail"`
	PersonalPhone string               `json:"personalPhone"`
	Avatar        string               `json:"avatar"`
	PermanentAddr location.AddressInfo `json:"permanentAddr"`
	CurrentAddr   location.AddressInfo `json:"currentAddr"`
	MaritalStatus string               `json:"maritalStatus"`
	IsResidential bool                 `json:"isResidential"`
}

type EmployeeModel struct {
	base.Model
	TenantID     string                      `json:"tenantId"`
	CompanyID    uuid.UUID                   `json:"companyId"`
	Company      *company.CompanyModel       `json:"company,omitempty"`
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
	GeneralInfo  GeneralInfo                 `json:"generalInfo"`
}

func (model EmployeeModel) DataMapper() *EmployeeDB {
	return &EmployeeDB{
		TenantID:     model.TenantID,
		CompanyID:    model.CompanyID,
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
		GeneralInfo:  model.GeneralInfo,
	}
}
