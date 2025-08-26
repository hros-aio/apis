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
	FirstName     string                `json:"firstName" validate:"required,isAlpha" example:"John"`
	MiddleName    string                `json:"middleName" validate:"isAlpha" example:"Doe"`
	LastName      string                `json:"lastName" validate:"required,isAlpha" example:"Smith"`
	Birth         string                `json:"birth" validate:"required,isDateString" example:"1990-01-01"`
	Gender        string                `json:"gender" validate:"required" example:"male"`
	WorkEmail     string                `json:"workEmail" validate:"required,isEmail" example:"john.doe@example.com"`
	WorkPhone     string                `json:"workPhone" example:"123-456-7890"`
	PersonalEmail string                `json:"personalEmail" example:"john.doe@gmail.com"`
	PersonalPhone string                `json:"personalPhone" example:"098-765-4321"`
	Avatar        string                `json:"avatar" example:"https://example.com/avatar.jpg"`
	PermanentAddr *location.AddressInfo `json:"permanentAddr" validate:"nested"`
	CurrentAddr   *location.AddressInfo `json:"currentAddr" validate:"nested"`
	MaritalStatus string                `json:"maritalStatus" example:"single"`
	IsResidential bool                  `json:"isResidential" example:"true"`
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
