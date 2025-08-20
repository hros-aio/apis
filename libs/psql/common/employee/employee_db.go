package employee

import (
	"time"

	"github.com/google/uuid"
	"github.com/tinh-tinh/sqlorm/v2"
)

type EmployeeDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantID     string     `gorm:"column:tenant_id;varchar(64);not null" json:"tenantId"`
	CompanyID    uuid.UUID  `gorm:"column:company_id;varchar(255);index:idx_employees_company_id;not null" json:"companyId"`
	UserID       uuid.UUID  `gorm:"column:user_id;varchar(255);not null" json:"userId"`
	DepartmentID uuid.UUID  `gorm:"column:department_id;varchar(255);not null" json:"departmentId"`
	LocationID   uuid.UUID  `gorm:"column:location_id;varchar(255);not null" json:"locationId"`
	GradeID      uuid.UUID  `gorm:"column:grade_id;varchar(255);not null" json:"gradeId"`
	TitleID      uuid.UUID  `gorm:"column:title_id;varchar(255);not null" json:"titleId"`
	FullName     string     `gorm:"column:full_name;type:varchar(255);not null" json:"fullName"`
	Code         string     `gorm:"column:code;type:varchar(255);not null" json:"code"`
	JoiningDate  time.Time  `gorm:"column:joining_date;type:timestamp;not null" json:"joiningDate"`
	LeftDate     *time.Time `gorm:"column:left_date;type:timestamp" json:"leftDate"`
	Type         string     `gorm:"column:type;type:varchar(255);not null" json:"type"`
	Status       string     `gorm:"column:status;type:varchar(255);not null" json:"status"`
}

func (EmployeeDB) TableName() string {
	return "employees"
}

func (data EmployeeDB) Dto() *EmployeeModel {
	return &EmployeeModel{
		TenantID:     data.TenantID,
		CompanyID:    data.CompanyID,
		UserID:       data.UserID,
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
	}
}
