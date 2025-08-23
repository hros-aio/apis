package company

import (
	"time"

	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
)

type CompanyDB struct {
	sqlorm.Model     `gorm:"embedded"`
	TenantID         string                 `gorm:"column:tenant_id;type:varchar(64);not null;index:idx_company_tenant_id" json:"tenantId"`
	Name             string                 `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	Legalname        string                 `gorm:"column:legal_name;type:varchar(255);not null;" json:"legalName"`
	RegistrationNo   string                 `gorm:"column:registration_no;type:varchar(100);not null;" json:"registrationNo"`
	TaxID            string                 `gorm:"column:tax_id;type:varchar(100);not null;" json:"taxId"`
	Website          string                 `gorm:"column:website;type:varchar(255);not null;" json:"website"`
	Industry         string                 `gorm:"column:industry;type:varchar(64);not null;" json:"industry"`
	Size             int                    `gorm:"column:size;type:int;not null;" json:"size"`
	Logo             string                 `gorm:"column:logo;type:varchar(256);not null;" json:"logo"`
	Contact          tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:contact_" json:"contact"`
	FoundedDate      time.Time              `gorm:"column:founded_date;type:date;not null;" json:"foundedDate"`
	SecondaryContact tenant.ContactPersonDb `gorm:"embedded;embeddedPrefix:secondary_contact_" json:"secondaryContact"`
}

func (CompanyDB) TableName() string {
	return "companies"
}

func (data CompanyDB) Dto() *CompanyModel {
	return &CompanyModel{
		Model:            base.Model(data.Model),
		TenantID:         data.TenantID,
		Name:             data.Name,
		Industry:         data.Industry,
		Size:             data.Size,
		Logo:             data.Logo,
		Contact:          tenant.ContactPerson(data.Contact),
		SecondaryContact: tenant.ContactPerson(data.SecondaryContact),
	}
}
