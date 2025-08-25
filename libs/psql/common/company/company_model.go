package company

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/tinh-tinh/sqlorm/v2"
)

type CompanyModel struct {
	base.Model
	TenantID         string               `json:"tenantId"`
	Name             string               `json:"name"`
	Legalname        string               `json:"legalName"`
	Status           string               `json:"status"`
	RegistrationNo   string               `json:"registrationNo"`
	TaxID            string               `json:"taxId"`
	Website          string               `json:"website"`
	Industry         string               `json:"industry"`
	Size             int                  `json:"size"`
	Logo             string               `json:"logo"`
	FoundedDate      time.Time            `json:"foundedDate"`
	Contact          tenant.ContactPerson `json:"contact"`
	SecondaryContact tenant.ContactPerson `json:"secondaryContact"`
	HoldingID        *uuid.UUID           `json:"holdingId,omitempty"`
	Holding          *CompanyModel        `json:"holding,omitempty"`
}

func (model CompanyModel) DataMapper() *CompanyDB {
	data := &CompanyDB{
		TenantID:         model.TenantID,
		Name:             model.Name,
		Legalname:        model.Legalname,
		Industry:         model.Industry,
		RegistrationNo:   model.RegistrationNo,
		TaxID:            model.TaxID,
		Website:          model.Website,
		FoundedDate:      model.FoundedDate,
		Status:           model.Status,
		Size:             model.Size,
		Logo:             model.Logo,
		Contact:          tenant.ContactPersonDb(model.Contact),
		SecondaryContact: tenant.ContactPersonDb(model.SecondaryContact),
		HoldingID:        model.HoldingID,
	}

	if !reflect.ValueOf(model.Model).IsZero() {
		data.Model = sqlorm.Model(model.Model)
	}

	return data
}
