package companies

import "github.com/hros-aio/apis/libs/sql/common/tenant"

type CreateCompanyInput struct {
	Name             string                `json:"name" example:"Terralogic Inc"`
	Industry         string                `json:"industry" example:"Techincal"`
	Size             int                   `json:"size" example:"50"`
	Logo             string                `json:"logo" example:"https://img.url"`
	Contact          *tenant.ContactPerson `json:"contact" validate:"nested"`
	SecondaryContact *tenant.ContactPerson `json:"secondaryContact" validate:"nested"`
}
