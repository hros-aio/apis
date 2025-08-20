package grades

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/tinh-tinh/sqlorm/v2"
)

type GradeDB struct {
	sqlorm.Model `gorm:"embedded"`
	TenantID     string             `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_location_tenant_id" json:"tenantId"`
	CompanyID    uuid.UUID          `gorm:"column:company_id" json:"companyId"`
	Company      *company.CompanyDB `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Code         string             `gorm:"column:code;type:varchar(64);not null" json:"code"`
	Name         string             `gorm:"column:name;type:varchar(255)" json:"name"`
}

func (GradeDB) TableName() string {
	return "grades"
}

func (data GradeDB) Dto() *GradeModel {
	model := &GradeModel{
		Model:     base.Model(data.Model),
		TenantID:  data.TenantID,
		CompanyID: data.CompanyID,
		Code:      data.Code,
		Name:      data.Name,
	}

	if data.Company != nil {
		model.Company = data.Company.Dto()
	}

	return model
}
