package companies

import (
	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/base"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/saga/messages"
)

func ToModel(data messages.CompanyActivatedPayload) *company.CompanyModel {
	return &company.CompanyModel{
		Model: base.Model{
			ID: uuid.MustParse(data.ID),
		},
		TenantID:       data.TenantID,
		Name:           data.Name,
		Legalname:      data.LegalName,
		Status:         data.Status,
		RegistrationNo: data.RegistrationNo,
		TaxID:          data.TaxID,
		Website:        data.Website,
		Industry:       data.Industry,
		Size:           data.Size,
		Logo:           data.Logo,
		FoundedDate:    data.FoundedDate,
		HoldingID:      data.HoldingID,
	}
}
