package companies

import (
	"time"

	"github.com/google/uuid"
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
)

type CreateCompanyInput struct {
	Name             string                `json:"name" example:"Abc Inc" validate:"required"`
	LegalName        string                `json:"legalName" example:"Abc Inc" validate:"required"`
	RegistrationNo   string                `json:"registrationNo" example:"123456789" validate:"required"`
	TaxID            string                `json:"taxId" example:"987654321" validate:"required"`
	Website          string                `json:"website" example:"https://abc.com" validate:"required"`
	FoundedDate      string                `json:"foundedDate" example:"2020-01-01" validate:"required"`
	Industry         string                `json:"industry" example:"Techincal" validate:"required"`
	Size             int                   `json:"size" example:"50" validate:"required"`
	Logo             string                `json:"logo" example:"https://img.url" validate:"required"`
	Contact          *tenant.ContactPerson `json:"contact" validate:"nested"`
	SecondaryContact *tenant.ContactPerson `json:"secondaryContact" validate:"nested"`
	HoldingID        *uuid.UUID            `json:"holdingId" validate:"isUUID"`
}

func (data *CreateCompanyInput) Dto() *company.CompanyModel {
	foundDate, _ := time.Parse("2006-01-02", data.FoundedDate)
	model := &company.CompanyModel{
		Name:           data.Name,
		Industry:       data.Industry,
		Size:           data.Size,
		Logo:           data.Logo,
		Legalname:      data.LegalName,
		RegistrationNo: data.RegistrationNo,
		TaxID:          data.TaxID,
		Website:        data.Website,
		HoldingID:      data.HoldingID,
		FoundedDate:    foundDate,
	}

	if data.Contact != nil {
		model.Contact = *data.Contact
	}

	if data.SecondaryContact != nil {
		model.SecondaryContact = *data.SecondaryContact
	}

	return model
}

type UpdateCompanyInput struct {
	Name             string                `json:"name,omitempty" example:"Terralogic Inc"`
	LegalName        string                `json:"legalName,omitempty" example:"Terralogic Inc"`
	RegistrationNo   string                `json:"registrationNo,omitempty" example:"123456789"`
	TaxID            string                `json:"taxId,omitempty" example:"987654321"`
	Website          string                `json:"website,omitempty" example:"https://img.url"`
	FoundedDate      string                `json:"foundedDate,omitempty" example:"2020-01-01"`
	Industry         string                `json:"industry,omitempty" example:"Techincal"`
	Size             int                   `json:"size,omitempty" example:"50"`
	Logo             string                `json:"logo,omitempty" example:"https://img.url"`
	Contact          *tenant.ContactPerson `json:"contact,omitempty" validate:"nested"`
	SecondaryContact *tenant.ContactPerson `json:"secondaryContact,omitempty" validate:"nested"`
	HoldingID        *uuid.UUID            `json:"holdingId,omitempty" validate:"isUUID"`
}

func (data *UpdateCompanyInput) Dto() *company.CompanyModel {
	model := &company.CompanyModel{
		Name:           data.Name,
		Industry:       data.Industry,
		Size:           data.Size,
		Logo:           data.Logo,
		Legalname:      data.LegalName,
		RegistrationNo: data.RegistrationNo,
		TaxID:          data.TaxID,
		Website:        data.Website,
		HoldingID:      data.HoldingID,
	}

	if data.FoundedDate != "" {
		foundDate, _ := time.Parse("2006-01-02", data.FoundedDate)
		model.FoundedDate = foundDate
	}

	if data.Contact != nil {
		model.Contact = *data.Contact
	}

	if data.SecondaryContact != nil {
		model.SecondaryContact = *data.SecondaryContact
	}

	return model
}
