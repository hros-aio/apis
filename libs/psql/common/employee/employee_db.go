package employee

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/department"
	"github.com/hros-aio/apis/libs/psql/common/location"
	"github.com/tinh-tinh/sqlorm/v2"
)

type EmployeeDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantID     string                   `gorm:"column:tenant_id;varchar(64);not null" json:"tenantId"`
	CompanyID    uuid.UUID                `gorm:"column:company_id;varchar(255);index:idx_employees_company_id;not null" json:"companyId"`
	Company      *company.CompanyDB       `gorm:"foreignKey:CompanyID" json:"company"`
	DepartmentID uuid.UUID                `gorm:"column:department_id;varchar(255);not null" json:"departmentId"`
	Department   *department.DepartmentDB `gorm:"foreignKey:DepartmentID" json:"department"`
	LocationID   uuid.UUID                `gorm:"column:location_id;varchar(255);not null" json:"locationId"`
	Location     *location.LocationDB     `gorm:"foreignKey:LocationID" json:"location"`
	GradeID      uuid.UUID                `gorm:"column:grade_id;varchar(255);not null" json:"gradeId"`
	TitleID      uuid.UUID                `gorm:"column:title_id;varchar(255);not null" json:"titleId"`
	FullName     string                   `gorm:"column:full_name;type:varchar(255);not null" json:"fullName"`
	Code         string                   `gorm:"column:code;type:varchar(255);not null" json:"code"`
	JoiningDate  time.Time                `gorm:"column:joining_date;type:timestamp;not null" json:"joiningDate"`
	LeftDate     *time.Time               `gorm:"column:left_date;type:timestamp" json:"leftDate"`
	Type         string                   `gorm:"column:type;type:varchar(255);not null" json:"type"`
	Status       string                   `gorm:"column:status;type:varchar(255);not null" json:"status"`
	GeneralInfo  GeneralInfo              `gorm:"column:general_info;type:jsonb" json:"generalInfo"`
}

func (EmployeeDB) TableName() string {
	return "employees"
}

func (data EmployeeDB) Dto() *EmployeeModel {
	model := &EmployeeModel{
		TenantID:     data.TenantID,
		CompanyID:    data.CompanyID,
		DepartmentID: data.DepartmentID,
		LocationID:   data.LocationID,
		GradeID:      data.GradeID,
		TitleID:      data.TitleID,
		FullName:     data.FullName,
		Code:         data.Code,
		JoiningDate:  data.JoiningDate,
		LeftDate:     data.LeftDate,
		Type:         data.Type,
		Status:       data.Status,
		GeneralInfo:  data.GeneralInfo,
	}

	if data.Company != nil {
		model.Company = data.Company.Dto()
	}

	if data.Department != nil {
		model.Department = data.Department.Dto()
	}

	if data.Location != nil {
		model.Location = data.Location.Dto()
	}

	return model
}

func (g GeneralInfo) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *GeneralInfo) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, g)
}
