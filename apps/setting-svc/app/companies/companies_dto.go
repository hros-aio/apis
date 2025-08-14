package companies

import (
	"github.com/hros-aio/apis/libs/psql/common/company"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
)

type CreateCompanyInput struct {
	Name             string                `json:"name" example:"Terralogic Inc" validate:"required"`
	Industry         string                `json:"industry" example:"Techincal" validate:"required"`
	Size             int                   `json:"size" example:"50" validate:"required"`
	Logo             string                `json:"logo" example:"https://img.url" validate:"required"`
	Contact          *tenant.ContactPerson `json:"contact" validate:"nested"`
	SecondaryContact *tenant.ContactPerson `json:"secondaryContact" validate:"nested"`
}

type UpdateCompanyInput struct {
	Name             string                `json:"name,omitempty" example:"Terralogic Inc"`
	Industry         string                `json:"industry,omitempty" example:"Techincal"`
	Size             int                   `json:"size,omitempty" example:"50"`
	Logo             string                `json:"logo,omitempty" example:"https://img.url"`
	Contact          *tenant.ContactPerson `json:"contact,omitempty" validate:"nested"`
	SecondaryContact *tenant.ContactPerson `json:"secondaryContact,omitempty" validate:"nested"`
}

func (data *UpdateCompanyInput) Dto() *company.CompanyModel {
	model := &company.CompanyModel{
		Name:     data.Name,
		Industry: data.Industry,
		Size:     data.Size,
		Logo:     data.Logo,
	}

	if data.Contact != nil {
		model.Contact = *data.Contact
	}

	if data.SecondaryContact != nil {
		model.SecondaryContact = *data.SecondaryContact
	}

	return model
}
